#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )"; cd ..; pwd -P )"
set -e

VERSION=$1
if [ -z "$1" ]; then
  VERSION="0.0.0"
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi


uname_m="$(uname -m)"
if [ "$uname_m" == "x86_64" ] || [ "$uname_m" == "amd64" ]; then
    ARCH=amd64
else
    echo "Machine architecture $uname_m is not supported for a test environment"
    exit 1
fi

random_string() {
    x=
    until [[ "$x" =~ [a-z] ]] && [[ "$x" =~ [A-Z] ]] && [[ "$x" =~ [0-9] ]]; do
        x=$( env LC_ALL=C dd if=/dev/urandom bs=256 count=1 2> /dev/null | LC_CTYPE=C tr -dc 'a-zA-Z0-9' | head -c 8 )
    done
    echo "$x"
}

test_result=1
tmpdir=$(mktemp -d -t soracom.XXXXXXXX)
trap 'tear_down' 0

tear_down() {
    : "Clean up tmpdir" && {
        [[ $tmpdir ]] && rm -rf "$tmpdir"
    }

    : "Report result" && {
        NO_COLOR=$'\e[0m'
        if [ "$test_result" -eq 0 ]; then
            GREEN=$'\e[0;32m'
            echo
            echo -e "${GREEN}TEST OK${NO_COLOR}"
            echo
        else
            RED=$'\e[0;31m'
            echo
            echo -e "${RED}TEST FAILED${NO_COLOR}"
            echo
        fi
        exit $test_result
    }
}

if [ -z "$SORACOM_AUTHKEY_ID_FOR_TEST" ] || [ -z "$SORACOM_AUTHKEY_FOR_TEST" ]; then
    echo
    echo "ERROR: Env vars SORACOM_AUTHKEY_ID_FOR_TEST and SORACOM_AUTHKEY_FOR_TEST are required to use the API sandbox."
    exit 1
fi

SORACOM_PROFILE_DIR=$tmpdir/.soracom
: "${SORACOM_ENDPOINT:=https://api-sandbox.soracom.io}"
SORACOM_ENVS=("SORACOM_ENDPOINT=$SORACOM_ENDPOINT" "SORACOM_PROFILE_DIR=$SORACOM_PROFILE_DIR" "SORACOM_DEBUG=$SORACOM_DEBUG")
EMAIL="soracom-cli-test+$(random_string)@soracom.jp"
PASSWORD=$(random_string)
ZIPCODE="1234567"
STATE="東京都"
CITY="港区"
ADDR1="赤坂1-2-3"
FULL_NAME="ソラコム 太郎"
PHONE="03-1234-5678"

OS=$( uname -s | tr '[:upper:]' '[:lower:]' )

SORACOM="$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}"

: "Create an account on the sandbox" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        configure-sandbox \
        --auth-key-id "$SORACOM_AUTHKEY_ID_FOR_TEST" \
        --auth-key "$SORACOM_AUTHKEY_FOR_TEST" \
        --email "$EMAIL" \
        --password "$PASSWORD" \
        --profile soracom-cli-test
}

: "Check if no subscribers are registered" && {
    n="$( env "${SORACOM_ENVS[@]}" "$SORACOM" subscribers list --profile soracom-cli-test | jq .[] | wc -l )"
    test "$n" -eq 0
}

: "Create a shipping address" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        shipping-addresses create \
        --zip-code "$ZIPCODE" \
        --state "$STATE" \
        --city "$CITY" \
        --address-line1 "$ADDR1" \
        --full-name "$FULL_NAME" \
        --phone-number "$PHONE" \
        --profile soracom-cli-test
        )"
    shipping_address_id="$( echo "$resp" | jq -r .shippingAddressId )"
}

: "Add coverage type 'g'" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        operator add-coverage-type \
        --coverage-type g \
        --profile soracom-cli-test
    )"
    echo "$resp"
}

: "Get current payment method for coverage type 'g'" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        payment-methods get-current \
        --coverage-type g \
        --profile soracom-cli-test
    )"
    echo "$resp"
}

: "Create an order: 3 SIM cards" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        orders create \
        --shipping-address-id "$shipping_address_id" \
        --body '{"orderItemList":[{"productCode":"4573326590013","quantity":3}]}' \
        --profile soracom-cli-test
        )"
    order_id="$( echo "$resp" | jq -r .orderId )"
}

: "Confirm the order" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        orders confirm \
        --order-id "$order_id" \
        --profile soracom-cli-test
        )"
}

: "🚢🇹Ship it!" && {
    resp="$( env SORACOM_VERBOSE=1 "${SORACOM_ENVS[@]}" "$SORACOM" \
        sandbox orders ship \
        --order-id "$order_id" \
        --profile soracom-cli-test
        )"
}

: "Check if subscribers are associated with the orer" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        orders list-subscribers \
        --order-id "$order_id" \
        --profile soracom-cli-test
        )"
    n="$( echo "$resp" | jq .orderedSubscriberList[].imsi | wc -l )"
    test "$n" -eq 3
}

: "Register the subscribers" && {
    resp="$( env SORACOM_VERBOSE=1 "${SORACOM_ENVS[@]}" "$SORACOM" \
        orders register-subscribers \
        --order-id "$order_id" \
        --profile soracom-cli-test
        )"
}

: "Check if the subscribers are registered (uses '--fetch-all' option to test pagination)" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers list \
        --fetch-all --limit 1 \
        --profile soracom-cli-test
        )"
    n="$( echo "$resp" | jq .[].imsi | wc -l )"
    test "$n" -eq 3
    imsi="$( echo "$resp" | jq -r .[0].imsi )"
}

: "Activate the first SIM" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers activate \
        --imsi "$imsi" \
        --profile soracom-cli-test
}

: "Change speed class to s1.fast" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers update-speed-class \
        --imsi "$imsi" \
        --speed-class "s1.fast" \
        --profile soracom-cli-test
}

: "Suspend the SIM" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers suspend \
        --imsi "$imsi" \
        --profile soracom-cli-test
}

: "Enable termination" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers enable-termination \
        --imsi "$imsi" \
        --profile soracom-cli-test
}

: "Terminate the SIM" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers terminate \
        --imsi "$imsi" \
        --profile soracom-cli-test
}

: "Check if the SIM is terminated" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers get \
        --imsi "$imsi" \
        --profile soracom-cli-test
        )"
    status="$( echo "$resp" | jq -r .status )"
    test "$status" = "terminated"
}

: "Create a SIM" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        sandbox subscribers create \
        --body '{"subscription":"plan-D"}' \
        --profile soracom-cli-test
        )"
    imsi="$( echo "$resp" | jq -r .imsi)"
    registrationSecret="$( echo "$resp" | jq -r .registrationSecret)"
}

: "Register the SIM" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers register \
        --imsi "$imsi" \
        --registration-secret "$registrationSecret" \
        --profile soracom-cli-test
}

: "Suspend the SIM" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers suspend \
        --imsi "$imsi" \
        --profile soracom-cli-test
}

: "Create a group" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        groups create \
        --body '{"tags":{"name":"test1"}}' \
        --profile soracom-cli-test
        )"
    groupId="$( echo "$resp" | jq -r .groupId)"
}

: "Put config to the group" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        groups put-config \
        --group-id "$groupId" \
        --namespace SoracomAir \
        --body '[{"key":"useVpg","value":true}]' \
        --profile soracom-cli-test
        )"
}

: "Sleep 15 seconds to make sure all subscribers indexed in the searchlight (elasticsearch)" && {
    sleep 15
}

: "Query subscribers" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        query subscribers \
        --imsi '00101' \
        --limit 10 \
        --profile soracom-cli-test
        )"
    numSubs="$( echo "$resp" | jq -r .[].imsi | wc -l )"
    test "$numSubs" -eq 4
}

: "Check if an error is returned when required parameter is missing" && {
    set +e
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers update-speed-class \
        --profile soracom-cli-test \
        2>&1 )"
    exitCode="$?"
    set -e
    test "$exitCode" -ne 0
    [[ "$resp" == *"Error: required parameter 'imsi' is not specified"* ]]
}

: "Check if an error is returned when required parameter in the request body is missing" && {
    set +e
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers update-speed-class \
        --imsi "001010000000000" \
        --profile soracom-cli-test \
        2>&1 )"
    exitCode="$?"
    set -e
    test "$exitCode" -ne 0
    [[ "$resp" == *"Error: required parameter 'speedClass' in body (or command line option 'speed-class') is not specified"* ]]
}

: "Checking english help text" && {
    help_en="$( env LC_ALL=en_US.UTF-8 "${SORACOM_ENVS[@]}" "$SORACOM" -h )"
    diff <( echo "$help_en" ) <( cat "$d/test/data/help_en_expected.txt" )
}

: "Checking japanese help text" && {
    help_ja=$( env LC_ALL=ja_JP.UTF-8 "${SORACOM_ENVS[@]}" "$SORACOM" -h )
    diff <( echo "$help_ja" ) <( cat "$d/test/data/help_ja_expected.txt" )
}

: "Displaying all top-level subcommands' help text"

: "Run soracom bills" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" bills --profile soracom-cli-test
    env "${SORACOM_ENVS[@]}" "$SORACOM" bills list --profile soracom-cli-test
}

: "Run soracom completion" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" completion
}

: "Run soracom coupons" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" coupons
}

: "Run soracom credentials" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" credentials
}

: "Run soracom event_handlers" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" event-handlers
}

: "Run soracom groups" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" groups
}

: "Run soracom operator" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" operator
}

: "Run soracom orders" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" orders
}

: "Run soracom payment_history" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" payment-history
}

: "Run soracom payment_methods" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" payment-methods
    env "${SORACOM_ENVS[@]}" "$SORACOM" payment-methods webpay
}

: "Run soracom products" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" products
}

: "Run soracom roles" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" roles
}

: "Run soracom shipping_addresses" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" shipping-addresses
}

: "Run soracom stats" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" stats
    env "${SORACOM_ENVS[@]}" "$SORACOM" stats air
    env "${SORACOM_ENVS[@]}" "$SORACOM" stats beam
}

: "Run soracom subscribers" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" subscribers
}

: "Run soracom users" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" users
    env "${SORACOM_ENVS[@]}" "$SORACOM" users auth-keys
    env "${SORACOM_ENVS[@]}" "$SORACOM" users password
    env "${SORACOM_ENVS[@]}" "$SORACOM" users permissions
}

: "Run soracom version" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" version
}

: "Run soracom vpg" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" vpg
}

test_result=0

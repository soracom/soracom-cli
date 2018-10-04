#!/bin/bash
d="$( cd "$( dirname "$0" )"; cd ..; pwd -P )"
set -e

VERSION=$1
if [ -z "$1" ]; then
  VERSION="0.0.0"
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi


uname_s="$(uname -s)"
if [ "$uname_s" == "Darwin" ]; then
    OS=darwin
elif [ "$uname_s" == "Linux" ]; then
    OS=linux
else
    echo "Operating system $uname_s is not supported for a test environment"
    exit 1
fi

uname_m="$(uname -m)"
if [ "$uname_m" == "x86_64" ]; then
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
            echo -e ${GREEN}'TEST OK'${NO_COLOR}
            echo
        else
            RED=$'\e[0;31m'
            echo
            echo -e ${RED}'TEST FAILED'${NO_COLOR}
            echo
        fi
        exit $test_result
    }
}

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

: "Extract binary" && {
    if [ "$OS" == "darwin" ]; then
        SORACOM="$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}"
    elif [ "$OS" == "linux" ]; then
        tar xvzf "$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}.tar.gz" -C "$d/soracom/dist/$VERSION"
        SORACOM="$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}/soracom"
    fi
}

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

: "Create an order" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        orders create \
        --shipping-address-id "$shipping_address_id" \
        --body '{"orderItemList":[{"productCode":"4573326590013","quantity":1}]}' \
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

: "Check if a subscriber is registered" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        subscribers list \
        --profile soracom-cli-test
        )"
    n="$( echo "$resp" | jq .[].imsi | wc -l )"
    test "$n" -eq 1
    imsi="$( echo "$resp" | jq -r .[].imsi )"
}

: "Activate the SIM" && {
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
        subscribers list \
        --profile soracom-cli-test
        )"
    status="$( echo "$resp" | jq -r .[].status )"
    test "$status" = "terminated"
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

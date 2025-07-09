#!/usr/bin/env bash
set -Eeuo pipefail
d="$( cd "$( dirname "$0" )"; cd ..; pwd -P )"

RESET=$'\e[0m'
BOLD=$'\e[1m'
GREEN=$'\e[0;32m'
RED=$'\e[0;31m'

VERSION=${1:-}
if [ -z "$VERSION" ]; then
  VERSION="0.0.0"
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi


uname_m="$(uname -m)"
if [ "$uname_m" == "x86_64" ] || [ "$uname_m" == "amd64" ]; then
    ARCH=amd64
elif [ "$uname_m" == "arm64" ]; then
    ARCH=arm64
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

progress() {
    local message=$1
    echo -n "${BOLD}${message}${RESET} " 1>&2
}

print_fail() {
    echo "${RED}FAIL${RESET}" 1>&2
}

print_ok() {
    echo "${GREEN}OK${RESET}" 1>&2
}

fail_with_message() {
    local message=$1

    print_fail
    echo -e "$message" 1>&2
    exit 1
}


if [ -z "$SORACOM_AUTHKEY_ID_FOR_TEST" ] || [ -z "$SORACOM_AUTHKEY_FOR_TEST" ]; then
    echo
    echo "ERROR: Env vars SORACOM_AUTHKEY_ID_FOR_TEST and SORACOM_AUTHKEY_FOR_TEST are required to use the API sandbox."
    exit 1
fi

SORACOM_PROFILE_DIR=$tmpdir/.soracom
: "${SORACOM_ENDPOINT:=https://api-sandbox.soracom.io}"
SORACOM_ENVS=("SORACOM_ENDPOINT=$SORACOM_ENDPOINT" "SORACOM_PROFILE_DIR=$SORACOM_PROFILE_DIR" "SORACOM_DEBUG=${SORACOM_DEBUG:-}")
EMAIL1="soracom-cli-test+$(random_string)@soracom.jp"
EMAIL2="soracom-cli-test+$(random_string)@soracom.jp"
EMAIL3="soracom-cli-test+$(random_string)@soracom.jp"
EMAIL4="soracom-cli-test+$(random_string)@soracom.jp"
PROFILE1="soracom-cli-test"
PROFILE2="soracom-cli-test2"
PROFILE3="soracom-cli-test3"
PROFILE4="soracom-cli-test4"
PASSWORD=$(random_string)
ZIPCODE="1234567"
STATE="東京都"
CITY="港区"
ADDR1="赤坂1-2-3"
FULL_NAME="ソラコム 太郎"
PHONE="03-1234-5678"

OS=$( uname -s | tr '[:upper:]' '[:lower:]' )

SORACOM="$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}"

invoke_soracom_command_without_profile() {
    env "${SORACOM_ENVS[@]}" "$SORACOM" "$@"
}

invoke_soracom_command_with_profile() {
    local profile=$1
    shift
    env "${SORACOM_ENVS[@]}" "$SORACOM" --profile "$profile" "$@"
}

invoke_soracom_command() {
    invoke_soracom_command_with_profile soracom-cli-test "$@"
}


#
# Start testing by creating operators on the sandbox
#

create_sandbox_account() {
    local email=$1
    local profile=$2

    progress "Creating an account with email $email on the sandbox"

    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        configure-sandbox \
        --auth-key-id "$SORACOM_AUTHKEY_ID_FOR_TEST" \
        --auth-key "$SORACOM_AUTHKEY_FOR_TEST" \
        --email "$email" \
        --password "$PASSWORD" \
        --coverage-type "jp" \
        --profile "$profile"
        )"

    print_ok
}

create_sandbox_account "$EMAIL1" "$PROFILE1"
create_sandbox_account "$EMAIL2" "$PROFILE2"
create_sandbox_account "$EMAIL3" "$PROFILE3"
create_sandbox_account "$EMAIL4" "$PROFILE4"

#
# Remember the operator IDs for each operator, to be used in the test cases later.
#

get_operator_id() {
    local profile=$1

    resp="$( invoke_soracom_command_with_profile "$profile" operator get 2>/dev/null )"

    echo "$resp" | jq -r .operatorId
}

opid1="$( get_operator_id "$PROFILE1" )"
opid2="$( get_operator_id "$PROFILE2" )"
opid3="$( get_operator_id "$PROFILE3" )"
opid4="$( get_operator_id "$PROFILE4" )"

#
# Create AuthKeys for each operator, and remember them.
#

auth_key_id_for_sandbox_user1=""
auth_key_id_for_sandbox_user2=""
auth_key_id_for_sandbox_user3=""
auth_key_id_for_sandbox_user4=""
auth_key_for_sandbox_user1=""
auth_key_for_sandbox_user2=""
auth_key_for_sandbox_user3=""
auth_key_for_sandbox_user4=""

create_auth_key_for_user1() {
    local resp

    progress "Generating authKeyId and authKey for the user1 on the sandbox"
    resp="$( invoke_soracom_command operator auth-keys generate 2>/dev/null )"

    auth_key_id_for_sandbox_user1="$( echo "$resp" | jq -r '.authKeyId' )"
    auth_key_for_sandbox_user1="$( echo "$resp" | jq -r '.authKey' )"

    print_ok
}

create_auth_key_for_user1

create_auth_key_for_user2() {
    local resp

    progress "Generating authKeyId and authKey for the user2 on the sandbox"
    resp="$( invoke_soracom_command_with_profile "$PROFILE2" operator auth-keys generate 2>/dev/null )"

    auth_key_id_for_sandbox_user2="$( echo "$resp" | jq -r '.authKeyId' )"
    auth_key_for_sandbox_user2="$( echo "$resp" | jq -r '.authKey' )"

    print_ok
}

create_auth_key_for_user2

create_auth_key_for_user3() {
    local resp

    progress "Generating authKeyId and authKey for the user3 on the sandbox"
    resp="$( invoke_soracom_command_with_profile "$PROFILE3" operator auth-keys generate 2>/dev/null )"

    auth_key_id_for_sandbox_user3="$( echo "$resp" | jq -r '.authKeyId' )"
    auth_key_for_sandbox_user3="$( echo "$resp" | jq -r '.authKey' )"

    print_ok
}

create_auth_key_for_user3

create_auth_key_for_user4() {
    local resp

    progress "Generating authKeyId and authKey for the user4 on the sandbox"
    resp="$( invoke_soracom_command_with_profile "$PROFILE4" operator auth-keys generate 2>/dev/null )"

    auth_key_id_for_sandbox_user4="$( echo "$resp" | jq -r '.authKeyId' )"
    auth_key_for_sandbox_user4="$( echo "$resp" | jq -r '.authKey' )"

    print_ok
}

create_auth_key_for_user4

#
# Get API Keys and Tokens for the generated users
#

api_key_for_sandbox_user1=""
#api_key_for_sandbox_user2=""
#api_key_for_sandbox_user3=""
api_key_for_sandbox_user4=""
api_token_for_sandbox_user1=""
#api_token_for_sandbox_user2=""
#api_token_for_sandbox_user3=""
api_token_for_sandbox_user4=""

get_api_key_and_token_for_user1() {
    local auth_result

    progress "Obtaining api key and token for the user1 on the sandbox"

    auth_result="$( invoke_soracom_command_without_profile auth --auth-key-id "$auth_key_id_for_sandbox_user1" --auth-key "$auth_key_for_sandbox_user1"  )"

    api_key_for_sandbox_user1="$( echo "$auth_result" | jq -r .apiKey )"
    api_token_for_sandbox_user1="$( echo "$auth_result" | jq -r .token )"

    print_ok
}

get_api_key_and_token_for_user1

#get_api_key_and_token_for_user2() {
#    local auth_result
#
#    progress "Obtaining api key and token for the user2 on the sandbox"
#
#    auth_result="$( invoke_soracom_command_without_profile auth --auth-key-id "$auth_key_id_for_sandbox_user2" --auth-key "$auth_key_for_sandbox_user2" 2>/dev/null )"
#
#    api_key_for_sandbox_user2="$( echo "$auth_result" | jq -r .apiKey )"
#    api_token_for_sandbox_user2="$( echo "$auth_result" | jq -r .token )"
#
#    print_ok
#}
#
#get_api_key_and_token_for_user2
#
#get_api_key_and_token_for_user3() {
#    local auth_result
#
#    progress "Obtaining api key and token for the user3 on the sandbox"
#
#    auth_result="$( invoke_soracom_command_without_profile auth --auth-key-id "$auth_key_id_for_sandbox_user3" --auth-key "$auth_key_for_sandbox_user3" 2>/dev/null )"
#
#    api_key_for_sandbox_user3="$( echo "$auth_result" | jq -r .apiKey )"
#    api_token_for_sandbox_user3="$( echo "$auth_result" | jq -r .token )"
#
#    print_ok
#}
#
#get_api_key_and_token_for_user3

get_api_key_and_token_for_user4() {
    local auth_result

    progress "Obtaining api key and token for the user4 on the sandbox"

    auth_result="$( invoke_soracom_command_without_profile auth --auth-key-id "$auth_key_id_for_sandbox_user4" --auth-key "$auth_key_for_sandbox_user4" 2>/dev/null )"

    api_key_for_sandbox_user4="$( echo "$auth_result" | jq -r .apiKey )"
    api_token_for_sandbox_user4="$( echo "$auth_result" | jq -r .token )"

    print_ok
}

get_api_key_and_token_for_user4


#
# Test if authentication by AuthKey is working.
#

try_to_authenticate_using_user1_auth_key_id_and_auth_key() {
    local resp
    local opid

    progress "Try to authenticate using the generated authKeyId and authKey"

    resp="$( invoke_soracom_command_without_profile operator get --auth-key-id "$auth_key_id_for_sandbox_user1" --auth-key "$auth_key_for_sandbox_user1" 2>/dev/null )"
    opid="$( echo "$resp" | jq -r .operatorId )"
    if [ "$opid1" != "$opid" ]; then
        fail_with_message "expected authenticated as $opid1, but: $opid"
    fi

    print_ok
}

try_to_authenticate_using_user1_auth_key_id_and_auth_key

#
# Test error cases for AuthKey authentication
#

try_to_authenticate_using_auth_key_id_only() {
    local resp
    local expected_error_message="Error: both --auth-key-id and --auth-key must be specified"

    progress "Try to authenticate using the generated authKeyId only (expecting an error)"

    set +e
    resp="$( invoke_soracom_command_without_profile operator get --auth-key-id "$auth_key_id_for_sandbox_user1" 2>&1 && fail_with_message "expected an error but not" )"
    set -e
    if ! [[ "$resp" =~ $expected_error_message ]]; then
        fail_with_message "unexpected error message: $resp"
    fi

    print_ok
}

try_to_authenticate_using_auth_key_id_only

try_to_authenticate_using_auth_key_only() {
    local resp
    local expected_error_message="Error: both --auth-key-id and --auth-key must be specified"

    progress "Try to authenticate using the generated authKey only (expecting an error)"

    set +e
    resp="$( invoke_soracom_command_without_profile operator get --auth-key "$auth_key_for_sandbox_user1" 2>&1 && fail_with_message "expected an error but not" )"
    set -e
    if ! [[ "$resp" =~ $expected_error_message ]]; then
        fail_with_message "unexpected error message: $resp"
    fi

    print_ok
}

try_to_authenticate_using_auth_key_only

#
# Test if authentication by AuthKey preceeds over authentication by profile command
#

test_if_auth_key_preceeds_over_profile_command() {
    local profile_command
    local resp
    local opid

    progress "Test if authKey preceeds over profile command"

    profile_command="printf {\"authKeyId\":\"$auth_key_id_for_sandbox_user1\",\"authKey\":\"$auth_key_for_sandbox_user1\"}"
    resp="$( invoke_soracom_command_without_profile operator get --profile-command "$profile_command" --auth-key-id "$auth_key_id_for_sandbox_user2" --auth-key "$auth_key_for_sandbox_user2" 2>/dev/null )"
    opid="$( echo "$resp" | jq -r .operatorId )"
    if [ "$opid2" != "$opid" ]; then
        fail_with_message "expected authenticated as $opid2, but: $opid"
    fi

    print_ok
}

test_if_auth_key_preceeds_over_profile_command

#
# Test if authentication by AuthKey preceeds over authentication by profile
#

test_if_auth_key_preceeds_over_profile() {
    local profile_command
    local resp
    local opid

    progress "Test if authKey preceeds over profile"

    profile_command="printf %s '{\"authKeyId\":\"$auth_key_id_for_sandbox_user2\",\"authKey\":\"$auth_key_for_sandbox_user2\"}'"
    resp="$( invoke_soracom_command_without_profile operator get --profile "$PROFILE1" --profile-command "$profile_command" --auth-key-id "$auth_key_id_for_sandbox_user3" --auth-key "$auth_key_for_sandbox_user3" 2>/dev/null )"
    opid="$( echo "$resp" | jq -r .operatorId )"
    if [ "$opid3" != "$opid" ]; then
        fail_with_message "expected authenticated as $opid3, but: $opid"
    fi

    print_ok
}

test_if_auth_key_preceeds_over_profile

#
# Test if authentication by profile command preceeds over authentication by profile
#

test_if_profile_command_preceeds_over_profile() {
    local profile_command
    local resp
    local opid

    progress "Test if profile command preceeds over profile"

    profile_command="printf %s '{\"authKeyId\":\"$auth_key_id_for_sandbox_user2\",\"authKey\":\"$auth_key_for_sandbox_user2\"}'"
    resp="$( invoke_soracom_command_without_profile operator get --profile "$PROFILE1" --profile-command "$profile_command" 2>/dev/null )"
    opid="$( echo "$resp" | jq -r .operatorId )"
    if [ "$opid2" != "$opid" ]; then
        fail_with_message "expected authenticated as $opid2 but: $opid"
    fi

    print_ok
}

test_if_profile_command_preceeds_over_profile

#
# Test if Profile Command specified in profile is working
#
test_profile_command_in_profile() {
    local profile_command
    local resp
    local opid

    local profile2="$SORACOM_PROFILE_DIR/soracom-cli-test2.json"

    progress "Test if profile command specified in profile is working"

    # inject user3's auth key into user2's profile's profileCommand
    profile_command="printf %s '{\\\"authKeyId\\\":\\\"$auth_key_id_for_sandbox_user3\\\",\\\"authKey\\\":\\\"$auth_key_for_sandbox_user3\\\"}'"
    mv "$profile2" "$profile2.bak"
    jq -r "(.profileCommand |= \"$profile_command\")" "$profile2.bak" > "$profile2"
    chmod 600 "$profile2"

    # authenticate with user2's profile, but it will be authenticated as user3, because profileCommand is preceeding
    resp="$( invoke_soracom_command_without_profile operator get --profile "$PROFILE2" 2>/dev/null )"
    opid="$( echo "$resp" | jq -r .operatorId )"
    if [ "$opid3" != "$opid" ]; then
        fail_with_message "expected authenticated as $opid3 but: $opid"
    fi

    mv "$profile2.bak" "$profile2"

    print_ok
}

test_profile_command_in_profile

#
# Test if API Key / Token preceeds all other authentication methods.
#
test_if_api_key_and_token_preceeds_all_other_authentication_methods() {
    local profile_command
    local resp
    local opid

    progress "Test if API Key and Token preceeds all other authentication methods"

    profile_command="printf %s '{\"authKeyId\":\"$auth_key_id_for_sandbox_user2\",\"authKey\":\"$auth_key_for_sandbox_user2\"}'"
    resp="$( invoke_soracom_command_without_profile operator get \
        --profile "$PROFILE1" \
        --profile-command "$profile_command" \
        --auth-key-id "$auth_key_id_for_sandbox_user3" --auth-key "$auth_key_for_sandbox_user3" \
        --api-key "$api_key_for_sandbox_user4" --api-token "$api_token_for_sandbox_user4" \
        2>/dev/null )"
    opid="$( echo "$resp" | jq -r .operatorId )"
    if [ "$opid4" != "$opid" ]; then
        fail_with_message "expected authenticated as $opid4, but: $opid"
    fi

    print_ok
}

test_if_api_key_and_token_preceeds_all_other_authentication_methods


#
# Test error cases for API Key and Token
#

try_to_call_api_only_with_api_key() {
    local resp
    local expected_error_message="Error: both --api-key and --api-token must be specified"

    progress "Try to call API using the API Key only (expecting an error)"

    set +e
    resp="$( invoke_soracom_command_without_profile operator get --api-key "$api_key_for_sandbox_user1" 2>&1 && fail_with_message "expected an error but not" )"
    set -e
    if ! [[ "$resp" =~ $expected_error_message ]]; then
        fail_with_message "unexpected error message: $resp"
    fi

    print_ok
}

try_to_call_api_only_with_api_key

try_to_call_api_only_with_api_token() {
    local resp
    local expected_error_message="Error: both --api-key and --api-token must be specified"

    progress "Try to call API using the API Token only (expecting an error)"

    set +e
    resp="$( invoke_soracom_command_without_profile operator get --api-token "$api_token_for_sandbox_user1" 2>&1 && fail_with_message "expected an error but not" )"
    set -e
    if ! [[ "$resp" =~ $expected_error_message ]]; then
        fail_with_message "unexpected error message: $resp"
    fi

    print_ok
}

try_to_call_api_only_with_api_token

try_to_call_api_without_any_credentials() {
    local resp
    local expected_error_message="Error: stat .*/\.soracom/default.json: no such file or directory"

    progress "Try to call API without any credentials (expecting an error)"

    set +e
    resp="$( invoke_soracom_command_without_profile operator get 2>&1 && fail_with_message "expected an error but not" )"
    set -e
    if ! [[ "$resp" =~ $expected_error_message ]]; then
        fail_with_message "unexpected error message: $resp"
    fi

    print_ok
}

try_to_call_api_without_any_credentials

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

: "Get subscribers list with --jsonl option" && {
    subscribers="$( invoke_soracom_command subscribers list --jsonl )"
    if [ "$( echo "$subscribers" | wc -l )" -ne 4 ]; then
        echo "expected 4 lines for 4 subscribers, but got the following:" 2>&1
        echo "$subscribers" 2>&1
        exit 1
    fi
}

groupName="test_group_name_$( random_string )"
: "Create a group" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        groups create \
        --body "{\"tags\":{\"name\":\"$groupName\"}}" \
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

: "Add a subscriber to the group" && {
    invoke_soracom_command subscribers set-group --imsi "$imsi" --group-id "$groupId"
}

: "Wait for all subscribers indexed in the searchlight (elasticsearch)" && {
    sleepSeconds=15
    echo "Sleeping $sleepSeconds seconds to make sure all subscribers are indexed in the database ..."
    for (( i=0; i < sleepSeconds; i++ )); do
        echo -n '.'
        sleep 1
    done
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

: "Query SIMs" && {
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        query sims \
        --imsi '00101' \
        --limit 10 \
        --profile soracom-cli-test
        )"
    numSubs="$( echo "$resp" | jq -r .[].simId | wc -l )"
    test "$numSubs" -eq 4
}

#: "Query subscribers by group name" && {
#    resp="$( invoke_soracom_command \
#        query sims \
#        --group "$groupName" \
#        --limit 10
#        )"
#    numSubs="$( echo "$resp" | jq -r .[].simId | wc -l )"
#    test "$numSubs" -eq 1
#}

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

: "Create Soralets" && {
    invoke_soracom_command soralets create --soralet-id soracom-cli-test
}

: "Upload wasm module to the Soralet" && {
    invoke_soracom_command soralets upload \
      --soralet-id soracom-cli-test \
      --content-type "application/octet-stream" \
      --body @"$d/test/data/gps-multi-unit.wasm"
}

: "Execute the wasm module" && {
    invoke_soracom_command soralets exec \
      --soralet-id soracom-cli-test \
      --version 1 \
      --direction uplink \
      --content-type "application/json" \
      --payload "{}" \
      --body '{"source":{"resourceType":"Subscriber","resourceId":"001010000000000"}}'
}

: "Delete the wasm module" && {
    invoke_soracom_command soralets delete-version --soralet-id soracom-cli-test --version 1
}

: "Delete the Soralet" && {
    invoke_soracom_command soralets delete --soralet-id soracom-cli-test
}

: "Checking english help text" && {
    help_en="$( env LC_ALL=en_US.UTF-8 "${SORACOM_ENVS[@]}" "$SORACOM" -h )"
    diff --ignore-trailing-space <( echo "$help_en" ) <( cat "$d/test/data/help_en_expected.txt" )
}

: "Checking japanese help text" && {
    help_ja=$( env LC_ALL=ja_JP.UTF-8 "${SORACOM_ENVS[@]}" "$SORACOM" -h )
    diff --ignore-trailing-space <( echo "$help_ja" ) <( cat "$d/test/data/help_ja_expected.txt" )
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

: "Should command execution fail when unnecessary arguments passed" && {
    set +e
    resp="$( env "${SORACOM_ENVS[@]}" "$SORACOM" \
        audit-logs api get \
        __UNNECESSARY__ __ARGUMENTS__ \
        --profile soracom-cli-test \
        2>&1 )"
    exitCode="$?"
    set -e
    test "$exitCode" -ne 0
    [[ "$resp" == *"Error: unexpected arguments passed => [__UNNECESSARY__ __ARGUMENTS__]"* ]]
}

test_result=0

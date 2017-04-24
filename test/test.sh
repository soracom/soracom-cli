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
        x=$( env LC_ALL=C tr -dc 'a-zA-Z0-9' < /dev/urandom | fold -w 8 | head -n 1)
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

: "Extract binary" && {
    if [ "$OS" == "darwin" ]; then
        SORACOM="$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}"
    elif [ "$OS" == "linux" ]; then
        tar xvzf "$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}.tar.gz" -C "$d/soracom/dist/$VERSION"
        SORACOM="$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}/soracom"
    fi
}

: "Create an account on the sandbox" && {
    go get -u github.com/soracom/soracom-sdk-go
    env SORACOM_ENDPOINT="$SORACOM_ENDPOINT" "${SORACOM_ENVS[@]}" go run "$d/test/setup.go" --email="$EMAIL" --password="$PASSWORD"
}

: "Checking english help text" && {
    help_en="$( env LC_ALL=en_US.UTF-8 "${SORACOM_ENVS[@]}" "$SORACOM" -h )"
    diff <( echo "$help_en" ) <( cat "$d/test/data/help_en_expected.txt" )
}

: "Checking japanese help text" && {
    help_ja=$( env LC_ALL=ja_JP.UTF-8 "${SORACOM_ENVS[@]}" "$SORACOM" -h )
    diff <( echo "$help_ja" ) <( cat "$d/test/data/help_ja_expected.txt" )
}

: "Run soracom configure and create the default profile" && {
    expect -c "$(cat <<EOD
spawn env ${SORACOM_ENVS[@]} $SORACOM configure
expect "\(1-2\)"
send -- "2\n"
expect "\(1-3\)"
send -- "2\n"
expect "email: "
send -- "$EMAIL\n"
expect "password: "
send -- "$PASSWORD\r\n"
expect eof
EOD
)"
}

: "Run soracom auth" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" auth --body "$( jo email="$EMAIL" password="$PASSWORD" )"
    env "${SORACOM_ENVS[@]}" "$SORACOM" auth --email="$EMAIL" --password="$PASSWORD"
}

: "Run soracom bills" && {
    env "${SORACOM_ENVS[@]}" "$SORACOM" bills
    env "${SORACOM_ENVS[@]}" "$SORACOM" bills list
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

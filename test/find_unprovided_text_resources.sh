#!/usr/bin/env bash
set -Eeuo pipefail
d="$( cd "$( dirname "$0" )"; cd ..; pwd -P )"

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

OS=$( uname -s | tr '[:upper:]' '[:lower:]' )

SORACOM="$d/soracom/dist/$VERSION/soracom_${VERSION}_${OS}_${ARCH}"
SORACOM_PROFILE_DIR=$tmpdir/.soracom
SORACOM_ENVS=("SORACOM_PROFILE_DIR=$SORACOM_PROFILE_DIR")

invoke_soracom_command() {
    env "${SORACOM_ENVS[@]}" sh -c "$SORACOM $*"
}

check_if_all_text_resources_provided() {
    local lang=$1
    shift
    local target_subcommand="$*"
    local child_subcommands

    echo -n '.'
    #echo "$target_subcommand"

    if ! all_text_resources_provided_for_subcommand "$lang" "$target_subcommand"; then
        echo
        echo "unprovided text resource found for 'LANG=$lang soracom $target_subcommand'"
        echo
        #exit 1
    fi

    child_subcommands="$( get_child_subcommands "$lang" "$target_subcommand" )"
    #echo "$child_subcommands"

    if [ -n "$child_subcommands" ]; then
        local child_subcommand
        while read -r child_subcommand; do
            if [ "$child_subcommand" == "help" ]; then
                continue
            fi
            #echo "$child_subcommand"
            check_if_all_text_resources_provided "$lang" "$target_subcommand" "$child_subcommand"
        done < <( echo "$child_subcommands" )
    fi
}

all_text_resources_provided_for_subcommand() {
    local lang=$1
    shift
    local target_subcommand="$*"
    local res
    res="$( LANG="$lang" invoke_soracom_command "$target_subcommand --help" </dev/null | grep 'cli\.'  2>&1 )"
    if [ -n "$res" ]; then
        return 1
    fi
    return 0
}

get_child_subcommands() {
    local lang=$1
    shift
    local target_subcommand="$*"
    res="$( LANG="$lang" invoke_soracom_command "$target_subcommand --help </dev/null" )"

    # https://stackoverflow.com/questions/38972736/how-to-print-lines-between-two-patterns-inclusive-or-exclusive-in-sed-awk-or/38972737#38972737
    echo "$res" | awk '/Available Commands:/{flag=1;next}/^$/{flag=0}flag { print $1 }'
}

echo "Checking English text resources"
check_if_all_text_resources_provided en
echo

echo "Checking Japanese text resources"
check_if_all_text_resources_provided ja
echo

test_result=0
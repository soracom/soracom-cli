#!/usr/bin/env bash
VERSION=$1
if [ -z "$1" ]; then
  VERSION='0.0.0'
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi

set -Eeuo pipefail

d="$( cd "$( dirname "$0" )" && cd .. && pwd -P )"
RED="\\033[1;31m"
GREEN="\\033[1;32m"
RESET="\\033[0m"

gopath=${GOPATH:-$HOME/go}
gopath=${gopath%%:*}

docker build -t soracom-cli-build "$d/build"
docker run --rm -t \
    --user "$(id -u):$(id -g)" \
    -e "VERSION=$VERSION" \
    -v "$d":/go/src/github.com/soracom/soracom-cli \
    -v "$gopath":/go \
    -v "$d/.cache":/.cache \
    -w "/go/src/github.com/soracom/soracom-cli/" \
    soracom-cli-build bash -x -c "/build/build.sh" || {
        echo
        echo -e "${RED}Build failed.${RESET}"
        echo
        exit 1
    }

echo
echo -e "${GREEN}OK${RESET}"
echo

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

mkdir -p "$d/.cache"
# Create the GOPATH beforehand; otherwise docker creates it as root when mounting
# it and the build fails with "permission denied".
mkdir -p "$gopath"

# go.mod is the single source of truth for the Go version to build with.
go_version="$( sed -n 's/^go \([0-9][0-9.]*\)$/\1/p' "$d/go.mod" )"
if [ -z "$go_version" ]; then
  echo -e "${RED}Unable to find the Go version in go.mod${RESET}"
  exit 1
fi

docker build -t soracom-cli-build "$d/build"
docker run --rm -t \
    --network host \
    --user "$(id -u):$(id -g)" \
    -e "VERSION=$VERSION" \
    -e "GOTOOLCHAIN=go$go_version" \
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

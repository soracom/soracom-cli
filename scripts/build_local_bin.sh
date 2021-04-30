#!/usr/bin/env bash
set -Eeuo pipefail
d="$( cd "$( dirname "$0" )" && cd .. && pwd -P )"

uname_s="$(uname -s)"
if [ "$uname_s" == "Darwin" ]; then
    OS=darwin
elif [ "$uname_s" == "Linux" ]; then
    OS=linux
elif [ "$uname_s" == "FreeBSD" ]; then
    OS=freebsd
else
    echo "Operating system $uname_s is not supported for a test environment"
    exit 1
fi

uname_m="$(uname -m)"
if [ "$uname_m" == "x86_64" ] || [ "$uname_m" == "amd64" ]; then
    ARCH=amd64
else
    echo "Machine architecture $uname_m is not supported for a test environment"
    exit 1
fi

VERSION=$1
if [ -z "$1" ]; then
  VERSION="0.0.0"
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi

pushd "$d/soracom" >/dev/null 2>&1
"$d/scripts/build.sh" "$VERSION" "$OS"
popd >/dev/null 2>&1

pushd "$d/soracom/dist/$VERSION/" >/dev/null 2>&1
if [ "$OS" == "darwin" ] || [ "$OS" == "freebsd" ]; then
    unzip -o soracom_${VERSION}_${OS}_${ARCH}.zip
elif [ "$OS" == "linux" ]; then
    tar xvzf soracom_${VERSION}_${OS}_${ARCH}.tar.gz
fi

popd >/dev/null 2>&1

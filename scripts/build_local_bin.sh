#!/bin/sh
d="$( cd "$( dirname "$0" )"; cd ..; pwd )"

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

pushd "$d/soracom" >/dev/null 2>&1
goxc -bc="$OS"
popd >/dev/null 2>&1

pushd "$GOPATH/bin/soracom-xc/snapshot/" >/dev/null 2>&1
if [ "$OS" == "darwin" ]; then
    unzip -fo soracom_${OS}_${ARCH}.zip
elif [ "$OS" == "linux" ]; then
    tar xvzf soracom_${OS}_${ARCH}.tar.gz
fi

popd >/dev/null 2>&1


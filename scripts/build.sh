#!/bin/bash
d=$( cd "$(dirname "$0" )"; cd ..; pwd )

: "Checking shell scripts" && {
    command -v shellcheck > /dev/null 2>&1
    if [ $? -eq 0 ]; then
        shellcheck scripts/*.sh
        shellcheck test/*.sh
    fi
}

: "Checking Prerequisites" && {
    command -v gox > /dev/null 2>&1
    if [ $? -ne 0 ]; then
        echo "gox must be installed"
        exit 1
    fi
}

set -e # aborting if any commands below exit with non-zero code

VERSION=$1
if [ -z "$1" ]; then
  VERSION="0.0.0"
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi

: "Installing dependencies" && {
    echo "Installing dependencies ..."
    go get -u golang.org/x/tools/cmd/goimports
    go get -u github.com/laher/goxc
    go get -u github.com/jteeuwen/go-bindata/...
    go get -u github.com/inconshreveable/mousetrap # required by spf13/cobra (only for windows env)
}

: "Testing generator's library" && {
    pushd "$d/generators/lib" > /dev/null
    go get ./...
    go test
    popd > /dev/null
}

: "Generating source code" && {
    echo "Generating command processor ..."
    pushd "$d/generators/cmd/src" > /dev/null
    go generate
    go get ./...
    go vet
    goimports -w ./*.go
    go test
    go build -o generate-cmd
    ./generate-cmd -a "$d/generators/assets/soracom-api.ftl.yaml" -t "$d/generators/cmd/templates" -p "$d/generators/cmd/predefined" -o "$d/soracom/generated/cmd/"
    popd > /dev/null
}

: "Building executables" && {
    pushd "$d/soracom" > /dev/null
    echo "Building artifacts ..."
    go generate
    go get ./...
    #gox -ldflags="-X github.com/soracom/soracom-cli/soracom/generated/cmd.version $VERSION" -osarch="windows/386 windows/amd64 darwin/amd64 linux/386 linux/amd64 linux/arm" -parallel=6 -output="bin/{{.OS}}/{{.Arch}}/soracom"
    goxc -bc="linux windows darwin" -d=dist/ -pv=$VERSION -build-ldflags="-X github.com/soracom/soracom-cli/soracom/generated/cmd.version=$VERSION"
    popd > /dev/null
}

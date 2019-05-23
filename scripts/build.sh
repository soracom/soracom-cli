#!/usr/bin/env bash

d=$( cd "$(dirname "$0" )"; cd ..; pwd -P )

I386="386"
ARM="arm"

: "Check if shell scripts are healthy" && {
  command -v shellcheck > /dev/null 2>&1 && {
    shellcheck -e SC2164 "$d/scripts/"*.sh
    shellcheck -e SC2164 "$d/test/"*.sh
  }
}

check_command_available() {
  local cmd=$1
  command -v "$cmd" > /dev/null 2>&1 || {
    echo "\`$cmd\` is required."
    exit 1
  }
}

convert_goarch_to_debian_arch() {
  local goarch=$1

  local debian_arch="$goarch"
  if [ "$goarch" = "$I386" ]; then
    debian_arch="i386"
  elif [ "$goarch" = "$ARM" ]; then
    debian_arch="armhf"
  fi

  echo $debian_arch
}

build_debian_control() {
  local version=$1
  local debian_arch=$2

  cat <<EOS
Package: soracom
Maintainer: SORACOM, INC.
Architecture: ${debian_arch}
Version: ${version}
Description: A command line tool \`soracom\` to invoke SORACOM API
EOS
}

: "Check if required commands for build are available" && {
  check_command_available go
  check_command_available git

  command -v dep > /dev/null 2>&1 || {
    go get -u github.com/golang/dep/cmd/dep
  }
}

set -e # aborting if any commands below exit with non-zero code
export GO111MODULE=on

VERSION=$1
if [ -z "$1" ]; then
  VERSION="0.0.0"
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi

LINUX="linux"
WINDOWS="windows"

TARGETS=$2
if [ -z "$2" ]; then
    TARGETS="$LINUX $WINDOWS darwin freebsd"
    uname_s=$( uname -s | tr '[:upper:]' '[:lower:]' )
    if [[ "$TARGETS" != *"$uname_s"* ]]; then
        TARGETS="$TARGETS $uname_s"
    fi
fi

ARCHTECTURES="amd64 $I386 $ARM"

# https://github.com/niemeyer/gopkg/issues/50
git config --global http.https://gopkg.in.followRedirects true

: "Install dependencies" && {
    echo "Installing build dependencies ..."
    go get -u golang.org/x/tools/cmd/goimports
    go get -u github.com/jessevdk/go-assets
    go get -u github.com/jessevdk/go-assets-builder
    go get -u github.com/elazarl/goproxy

    echo "Installing runtime dependencies ..."
    go get -u github.com/inconshreveable/mousetrap # required by spf13/cobra (only for windows env)
}

: "Test generator's library" && {
    pushd "$d/generators/lib" > /dev/null
    go test
    popd > /dev/null
}

: "Generate source code for soracom-cli" && {
    echo "Generating generator ..."
    pushd "$d/generators/cmd/src" > /dev/null
    go generate
    go vet
    goimports -w ./*.go
    go test
    go build -o generate-cmd

    echo "Generating source codes for soracom-cli by using the generator ..."
    ./generate-cmd -a "$d/generators/assets/soracom-api.en.yaml" -s "$d/generators/assets/sandbox/soracom-sandbox-api.en.yaml" -t "$d/generators/cmd/templates" -p "$d/generators/cmd/predefined" -o "$d/soracom/generated/cmd/"
    popd > /dev/null
}

: "Build soracom-cli executables" && {
    pushd "$d/soracom" > /dev/null
    echo "Building artifacts ..."
    go generate
    go get -u github.com/bearmini/go-acl # required to specify some dependencies explicitly as they are imported only in windows builds
    GOOS="$WINDOWS" go get -u golang.org/x/sys/windows
    gofmt -s -w .

    mkdir -p "dist/$VERSION/.tmp"

    for GOOS in $TARGETS; do
      for GOARCH in $ARCHTECTURES; do
        IDENT="soracom_${VERSION}_${GOOS}_${GOARCH}"
        BIN_FILENAME=$IDENT
        if [ "$GOOS" = "$WINDOWS" ]; then
          BIN_FILENAME="${IDENT}.exe"
        fi

        go build -ldflags "-X github.com/soracom/soracom-cli/soracom/generated/cmd.version=$VERSION" \
          -o "dist/${VERSION}/${BIN_FILENAME}"
        chmod 0755 "dist/${VERSION}/${BIN_FILENAME}"

        mkdir -p "dist/${VERSION}/.tmp/${IDENT}"
        cp "dist/${VERSION}/${BIN_FILENAME}" "dist/${VERSION}/.tmp/${IDENT}/soracom"

        (
          cd "dist/${VERSION}/.tmp"

          if [ "$GOOS" = "$LINUX" ]; then
            tar zcf "${IDENT}.tar.gz" "$IDENT"
            mv "${IDENT}.tar.gz" ../

            mkdir -p usr/bin/ DEBIAN
            cp "${IDENT}/soracom" usr/bin/soracom

            DEBIAN_ARCH=$(convert_goarch_to_debian_arch "$GOARCH")
            build_debian_control "$VERSION" "$DEBIAN_ARCH" > DEBIAN/control
            dpkg-deb --build . .
            mv "soracom_${VERSION}_${DEBIAN_ARCH}.deb" ../
          else
            zip -r "${IDENT}.zip" "$IDENT"
            mv "${IDENT}.zip" ../
          fi
        )
      done
    done

    rm -rf "dist/$VERSION/.tmp"

    # removing empty directories
    find "dist/$VERSION" -type d -empty -delete

    popd > /dev/null
}

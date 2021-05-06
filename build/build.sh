#!/usr/bin/env bash
if [ -z "$VERSION" ]; then
  VERSION='0.0.0'
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi

set -Eeuo pipefail

d="/go/src/github.com/soracom/soracom-cli"

: 'Check if shell scripts are healthy' && {
  shellcheck -e SC2164 "$d/scripts/"*.sh
  shellcheck -e SC2164 "$d/test/"*.sh
}

gopath=${GOPATH:-$HOME/go}
gopath=${gopath%%:*}

: 'Install dependencies' && {
    echo 'Installing build dependencies ...'
    go get -u golang.org/x/tools/cmd/goimports

    echo 'Installing commands used with "go generate" ...'
    go get -u github.com/elazarl/goproxy
    go mod tidy
}

: "Test generator's library" && {
    echo "Testing generator's source ..."
    go test "$d/generators/cmd/src"
    go test "$d/generators/lib"
}

: 'Generate source code for soracom-cli' && {
    echo 'Generating generator ...'
    (
      cd "$d/generators/cmd/src" && \
      go generate && \
      go vet && \
      goimports -w ./*.go && \
      go build -o generate-cmd
    )

    echo 'Generating source codes for soracom-cli by using the generator ...'
    "$d/generators/cmd/src/generate-cmd" -a generators/assets/soracom-api.en.yaml -s generators/assets/sandbox/soracom-sandbox-api.en.yaml -t generators/cmd/templates -p generators/cmd/predefined -o soracom/generated/cmd/

    echo 'Copying assets to embed ...'
    cp -r generators/assets/ soracom/generated/cmd/assets/
}

remove_tmpdir() {
  [[ -n "${1-}" ]] && rm -rf "$1"
}

build_deb_package() {
  goos=$1
  goarch=$2
  bindir=$3
  bin=$4

  case "$goarch" in
    "amd64" | "arm64")
      arch="$goarch"
      ;;
    "386")
      arch="i386"
      ;;
    "arm")
      arch="armhf"
      ;;
    *)
      arch="unknown"
      ;;
  esac

  fsize="$( stat --printf="%s" "$bindir/$bin" )"
  fsize="$(( fsize / 1024 ))"

  tmpdir="$( mktemp -d )"
  trap 'remove_tmpdir $tmpdir' RETURN

  mkdir -p "$tmpdir/DEBIAN"
  cat << EOD > "$tmpdir/DEBIAN/control"
Package: soracom
Priority: extra
Maintainer: soracom-cli@soracom.jp
Section: devel
Version: $VERSION
Architecture: $arch
Description: A command line tool \`soracom\' to invoke SORACOM API.
Installed-Size: $fsize
EOD

  mkdir -p "$tmpdir/usr/bin/"
  cp "$bindir/$bin" "$tmpdir/usr/bin/soracom"

  fakeroot dpkg-deb --build "$tmpdir" "$bindir/soracom_${VERSION}_$arch.deb" &> /dev/null
}

build() {
  goos=$1
  goarch=$2
  ext=$3

  bindir="soracom/dist/$VERSION"
  bin="soracom_${VERSION}_$goos-$goarch$ext"

  printf "  %-7s - %-5s  bin" "$goos" "$goarch"
  GOOS="$goos" GOARCH="$goarch" go build -ldflags="-X 'github.com/soracom/soracom-cli/soracom/generated/cmd.version=$VERSION'" -o "$bindir/$bin" ./soracom

  printf ", archive"
  case "$goos" in
    "linux" | "freebsd")
      tar -zcf "$bindir/$bin.tar.gz" "$bindir/$bin"
      ;;
    "darwin" | "windows")
      zip -q "$bindir/$bin.zip" "$bindir/$bin"
      ;;
    "*")
      echo "unknown GOOS. skipping compression"
      ;;
  esac

  if [ "$goos" == "linux" ]; then
    printf ", package"
    build_deb_package "$goos" "$goarch" "$bindir" "$bin"
  fi

  printf "\n"
}

: 'Build soracom-cli executables' && {
    echo 'Building artifacts ...'
    (
      cd "$d/soracom" && \
      go generate
      go get -u github.com/bearmini/go-acl # required to specify some dependencies explicitly as they are imported only in windows builds
      gofmt -s -w .
    )

    rm -rf "soracom/dist/$VERSION"
    build linux   amd64 ''
    build linux   arm64 ''
    build linux   386   ''
    build linux   arm   ''
    build darwin  amd64 ''
    build darwin  arm64 ''
    build freebsd amd64 ''
    build freebsd 386   ''
    build freebsd arm   ''
    build windows amd64 .exe
    build windows 386   .exe

}

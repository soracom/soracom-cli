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

: "Install dev dependencies" && {
  make install-dev-deps
}

: "Test generator's library" && {
  make test
}

: 'Generate source code for soracom-cli' && {
  make generate
}

: "Test the generated source code" && {
  make test-generated
  make lint
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
    "riscv64")
      arch="riscv64"
      ;;
    *)
      arch="unknown"
      ;;
  esac

  tmpdir="$( mktemp -d )"
  trap 'remove_tmpdir $tmpdir' RETURN

  mkdir -p "$tmpdir/usr/bin/"
  cp "$bindir/$bin" "$tmpdir/usr/bin/soracom"

  local license_text
  license_text="$( sed -e '1,4d' -e 's/^$/./' -e 's/^/ /' -e '$s/^ \.$//' "$d/LICENSE" )"

  mkdir -p "$tmpdir/usr/share/doc/soracom/"
  cat << EOD > "$tmpdir/usr/share/doc/soracom/copyright"
Format: https://www.debian.org/doc/packaging-manuals/copyright-format/1.0/
Upstream-Name: soracom-cli
Source: https://github.com/soracom/soracom-cli

Files: *
Copyright: 2015 Soracom, Inc.
License: MIT

License: MIT
${license_text}
EOD


  fsize="$( find $tmpdir -type f -not -path "$tmpdir/DEBIAN/*" -exec du -cb {} + | tail -1 | cut -f1 )"
  fsize="$(( fsize / 1024 ))"

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

  fakeroot dpkg-deb --build "$tmpdir" "$bindir/soracom_${VERSION}_$arch.deb" &> /dev/null
}

archive() {
  goos=$1
  goarch=$2
  bindir=$3
  binbase=$4
  ext=$5

  tmpdir="$( mktemp -d )"
  trap 'remove_tmpdir $tmpdir' RETURN

  workdir="soracom_${VERSION}_${goos}_$goarch"
  mkdir -p "$tmpdir/$workdir"
  cp "$bindir/$bin" "$tmpdir/$workdir/soracom$ext"

  case "$goos" in
    "linux" | "freebsd")
      tar -C "$tmpdir" -zcf "$bindir/$binbase.tar.gz" "$workdir/soracom$ext"
      ;;
    "darwin" | "windows")
      (cd "$tmpdir" && zip -q "$binbase.zip" "$workdir/soracom$ext")
      mv "$tmpdir/$binbase.zip" "$bindir/"
      ;;
    "*")
      echo "unknown GOOS. skipping compression"
      ;;
  esac
}

build() {
  goos=$1
  goarch=$2
  ext=$3

  bindir="soracom/dist/$VERSION"
  binbase="soracom_${VERSION}_${goos}_$goarch"
  bin="$binbase$ext"

  printf "  %-7s - %-7s  bin" "$goos" "$goarch"
  if [ "$goos" == "windows" ]; then
    make winres
  fi
  make build GOOS="$goos" GOARCH="$goarch" VERSION="$VERSION" OUTPUT="$bindir/$bin"

  printf ", archive"
  archive "$goos" "$goarch" "$bindir" "$binbase" "$ext"

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
    build linux   amd64   ''
    build linux   arm64   ''
    build linux   386     ''
    build linux   arm     ''
    build linux   riscv64 ''
    build darwin  amd64   ''
    build darwin  arm64   ''
    build freebsd amd64   ''
    build freebsd 386     ''
    build freebsd arm     ''
    build windows amd64   .exe
    build windows 386     .exe

}

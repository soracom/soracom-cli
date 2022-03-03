#!/usr/bin/env bash
set -Eeuo pipefail

if \soracom > /dev/null 2>&1; then
  if [[ "$( \command -v soracom )" != "/usr/local/bin/soracom" ]] || [ -L "/usr/local/bin/soracom" ]; then
    echo 'soracom-cli is already installed by using another method (brew, snap, apt etc.).' 2>&1
    echo 'Please use the same method if you want to update soracom-cli.' 2>&1
    exit 1
  fi
fi

if ! \curl --version > /dev/null 2>&1; then
  echo '"curl" command is required to install soracom-cli' 2>&1
  echo 'Please install "curl" command before proceeding.' 2>&1
  exit
fi

get_goos() {
  case "$( uname -s )" in
    "Linux")
      echo "linux"
      ;;
    "Darwin")
      echo "darwin"
      ;;
    "FreeBSD")
      echo "freebsd"
      ;;
    "Windows"*)
      echo "windows"
      ;;
    *)
      echo "unknown system: $( uname -s )" 2>&1
      exit 1
      ;;
  esac
}

get_goarch() {
  case "$( uname -m )" in
    "amd64" | "x86_64")
      echo "amd64"
      ;;
    "i386" | "i686")
      echo "386"
      ;;
    "armv"*)
      echo "arm"
      ;;
    "arm64" | "aarch64")
      echo "arm64"
      ;;
    *)
      echo "unknown architecture: $( uname -m )" 2>&1
      exit 1
      ;;
  esac
}

get_ext_regexp() {
  local goos=$1
  case "$goos" in
    "linux" | "freebsd")
      echo "\.tar\.gz"
      ;;
    "darwin" | "windows")
      echo "\.zip"
      ;;
    *)
      echo "unknown goos: $goos" 2>&1
      exit 1
      ;;
  esac
}

get_ext() {
  local goos=$1
  case "$goos" in
    "linux" | "freebsd")
      echo ".tar.gz"
      ;;
    "darwin" | "windows")
      echo ".zip"
      ;;
    *)
      echo "unknown goos: $goos" 2>&1
      exit 1
      ;;
  esac
}

extract() {
  local path=$1
  local dir=$2
  local ext=$3
  case "$ext" in
    ".tar.gz")
      tar -C "$dir" -xf "$path"
      ;;
    ".zip")
      unzip -q "$path" -d "$dir"
      ;;
    *)
      echo "unknown archive extension: $ext"
      exit 1
      ;;
  esac
}

goos="$( get_goos )"
goarch="$( get_goarch )"
ext_regexp="$( get_ext_regexp "$goos" )"
ext="$( get_ext "$goos" )"

url="$( \curl -fsSL https://api.github.com/repos/soracom/soracom-cli/releases/latest | \
  grep 'browser_download_url' | \
  grep "${goos}_${goarch}.*${ext_regexp}" | \
  cut -d : -f 2-3 | \
  tr -d '"'
)"
url="$(echo $url)" # trim spaces
fname=${url##*/}   # removes longest matching series of the pattern from the front (string after the last / will be left)

tmpdir=$(mktemp -d -t soracom.XXXXXXXX)
#trap 'tear_down' 0

tear_down() {
    : "Clean up tmpdir" && {
        [[ $tmpdir ]] && rm -rf "$tmpdir"
    }
}

echo -n "Downloading ... "
curl -fsSL --output "$tmpdir/$fname" "$url"
echo "done."

echo -n "Extracting ... "
extract "$tmpdir/$fname" "$tmpdir" "$ext"
echo "done."

echo -n "Installing ... "
dirname="${fname%"${ext}"}"
sudo mv "$tmpdir/$dirname/soracom" /usr/local/bin
sudo chmod +x /usr/local/bin/soracom
echo "done."

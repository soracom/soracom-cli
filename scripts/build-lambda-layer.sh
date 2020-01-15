#!/usr/bin/env bash
d=$( cd "$( dirname "$0" )"; cd ..; pwd -P )
set -e
set -x

VERSION=$1
if [ -z "$1" ]; then
  echo "Please specify version number (e.g. 1.2.3)"
  exit 1
fi

: "Generate layer.zip" && {
  wd="$d/soracom/dist/$VERSION/lambda-layer"
  rm -rf "$wd"
  mkdir -p "$wd/bin"
  curl -s -L "https://github.com/soracom/soracom-cli/releases/download/v${VERSION}/soracom_${VERSION}_linux_amd64" -o "$wd/bin/soracom"
  chmod +x "$wd/bin/soracom"
  cd "$wd" && zip -r "layer_${VERSION}.zip" . && cd -
}

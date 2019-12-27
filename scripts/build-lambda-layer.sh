#!/usr/bin/env bash
d=$( cd "$( dirname "$0" )"; cd ..; pwd -P )
set -e

VERSION=$1
if [ -z "$1" ]; then
  echo "Please specify version number (e.g. 1.2.3)"
  exit 1
fi

: "Generate layer.zip" && {
  vd="$d/soracom/dist/$VERSION"
  wd="$vd/lambda-layer"
  mkdir -p "$wd/bin"
  cp "$vd/soracom_${VERSION}_linux_amd64" "$wd/bin"
  cd "$wd" && zip -r "layer_${VERSION}.zip" . && cd -
}

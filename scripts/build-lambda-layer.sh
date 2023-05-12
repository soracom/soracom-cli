#!/usr/bin/env bash
set -Eeuo pipefail
d=$( cd "$( dirname "$0" )" && cd .. && pwd -P )

VERSION=${1:-}
if [ -z "$VERSION" ]; then
  echo "Please specify version number (e.g. 1.2.3)"
  exit 1
fi

generate_layer_zip() {
  wd=$1
  arch=$2
  curl -s -L "https://github.com/soracom/soracom-cli/releases/download/v${VERSION}/soracom_${VERSION}_linux_${arch}" -o "$wd/bin/soracom"
  chmod +x "$wd/bin/soracom"
  cd "$wd" && zip -r "layer_${VERSION}_${arch}.zip" . && cd -
}

wd="$d/soracom/dist/$VERSION/lambda-layer"
rm -rf "$wd"
mkdir -p "$wd/bin"

set -x
generate_layer_zip "$wd" amd64
generate_layer_zip "$wd" arm64
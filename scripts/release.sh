#!/usr/bin/env bash
VERSION=$1
if [ -z "$1" ]; then
  echo "Version number (e.g. 1.2.3) must be specified. Abort."
  exit 1
fi

set -Eeuo pipefail
d="$( cd "$( dirname "$0" )" && cd .. && pwd )"

command -v ghr > /dev/null 2>&1 || {
    echo "'ghr' is required."
    echo "Install ghr by following the instructions: https://github.com/tcnksm/ghr#install"
    exit 1
}

pushd "$d/soracom" >/dev/null 2>&1
rm -f "$d/soracom/dist/$VERSION/downloads.md"
ghr --prerelease --replace -u soracom -r soracom-cli "v$VERSION" "$d/soracom/dist/$VERSION/"
popd >/dev/null 2>&1

echo
echo "Please run \`update-homebrew-formula.sh\` as soon as the release gets published"
echo


#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )"; cd ..; pwd )"
set -e

VERSION=$1
if [ -z "$1" ]; then
  echo "Version number (e.g. 1.2.3) must be specified. Abort."
  exit 1
fi

pushd "$d/soracom" >/dev/null 2>&1
rm -f "$d/soracom/dist/$VERSION/downloads.md"
ghr --prerelease --replace -u soracom -r soracom-cli "v$VERSION" "$d/soracom/dist/$VERSION/"
popd >/dev/null 2>&1

echo
echo "Please run \`update-homebrew-formula.sh\` as soon as the release gets published"
echo


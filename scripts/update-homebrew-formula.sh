#!/bin/bash
set -e

version=$1
if [ "$version" == "" ]; then
    echo "Please specify version number (e.g. '1.2.3')"
    exit 1
fi

tmpdir=$(mktemp -d)

function cleanup {
    rm -rf "$tmpdir"
}

trap cleanup EXIT

cd "$tmpdir"
git clone git@github.com:soracom/homebrew-soracom-cli
cd homebrew-soracom-cli
git checkout master
bash ./update-formula.sh "$version"
git add soracom-cli.rb
git commit -m "bump version"
git push origin master

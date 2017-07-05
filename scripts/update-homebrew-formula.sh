#!/bin/bash
set -e

version=$1
git_username=$2
git_email=$3

if [[ "$version" == "" ]] || [[ "$git_username" == "" ]] || [[ "$git_email" == "" ]]; then
    echo "usage: $0 <version number (e.g. '1.2.3')> <git user.name> <git user.email>"
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
git config user.name "$git_username"
git config user.email "$git_email"
git commit -m "bump version"
git push origin master

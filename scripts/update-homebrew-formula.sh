#!/usr/bin/env bash
version=$1
git_username=$2
git_email=$3

if [[ "$version" == "" ]] || [[ "$git_username" == "" ]] || [[ "$git_email" == "" ]]; then
    echo "usage: $0 <version number (e.g. '1.2.3')> <git user.name> <git user.email>"
    exit 1
fi

# Check if GitHub CLI(gh) is installed
command -v gh > /dev/null 2>&1 || {
    echo "'gh' is required."
    echo "Install gh by following the instructions: https://github.com/cli/cli#installation"
    exit 1
}

set -Eeuo pipefail

tmpdir=$(mktemp -d)

function cleanup {
    rm -rf "$tmpdir"
}

trap cleanup EXIT

cd "$tmpdir"
git clone git@github.com:soracom/homebrew-soracom-cli
cd homebrew-soracom-cli
git checkout master

# Create feature branch
branch_name="bump-version-$version"
git checkout -b "$branch_name"

bash ./update-formula.sh "$version"
git add soracom-cli.rb
git config user.name "$git_username"
git config user.email "$git_email"
git commit -m "bump version to $version"

# Push feature branch and create PR (force push to overwrite existing branch)
git push --force-with-lease origin "$branch_name"
gh pr create --title "Bump version to $version" \
             --body "Bump version to $version"

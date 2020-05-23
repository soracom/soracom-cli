#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"

fen="$d/../soracom-api/build/soracom-api.en.yaml"
fja="$d/../soracom-api/build/soracom-api.ja.yaml"

if [ ! -f "$fen" ] || [ ! -f "$fja" ]; then
  echo "API definition files not found. Build soracom-api first."
  exit 1
fi

set -x
cp "$d/../soracom-api/build/soracom-api.en.yaml" "$d/generators/assets/"
cp "$d/../soracom-api/build/soracom-api.ja.yaml" "$d/generators/assets/"

cp "$d/../soracom-api/apidef/sandbox/soracom-sandbox-api.en.yaml" "$d/generators/assets/sandbox/"
cp "$d/../soracom-api/build/soracom-sandbox-api.ja.yaml"          "$d/generators/assets/sandbox/"
set +x

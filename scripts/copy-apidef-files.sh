#!/usr/bin/env bash
set -Eeuo pipefail
d="$( cd "$( dirname "$0" )" && cd .. && pwd )"

fen="$d/../soracom-api/build/soracom-api.en.yaml"
fja="$d/../soracom-api/build/soracom-api.ja.yaml"
sen="$d/../soracom-api/apidef/sandbox/soracom-sandbox-api.en.yaml"
sja="$d/../soracom-api/build/soracom-sandbox-api.ja.yaml"

if [ ! -f "$fen" ] || [ ! -f "$fja" ] || [ ! -f "$sen" ] || [ ! -f "$sja" ]; then
  echo "API definition files not found. Build soracom-api first."
  exit 1
fi

set -x
cp "$fen" "$d/generators/assets/"
cp "$fja" "$d/generators/assets/"
cp "$sen" "$d/generators/assets/sandbox/"
cp "$sja" "$d/generators/assets/sandbox/"
set +x

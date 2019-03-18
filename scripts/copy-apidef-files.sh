#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"

fen="$d/../soracom-api/build/soracom-api.en.yaml"
fja="$d/../soracom-api/build/soracom-api.ja.yaml"

if [ ! -f "$fen" ] || [ ! -f "$fja" ]; then
  echo "API definition files not found. Build soracom-api first."
  exit 1
fi

cp "$d/../soracom-api/build/soracom-api.en.yaml" "$d/generators/assets/"
cp "$d/../soracom-api/build/soracom-api.ja.yaml" "$d/generators/assets/"


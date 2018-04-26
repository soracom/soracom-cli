#!/bin/bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"

cp "$GOPATH/src/github.com/soracom/soracom-api/apidef/prod/soracom-api.en.yaml" "$d/generators/assets/"
cp "$GOPATH/src/github.com/soracom/soracom-api/apidef/prod/i18n/soracom-api.ja.yaml" "$d/generators/assets/"


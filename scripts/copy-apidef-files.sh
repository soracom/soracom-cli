#!/usr/bin/env bash
set -Eeuo pipefail
d="$( cd "$( dirname "$0" )" && cd .. && pwd )"

fen="$d/../soracom-api/dist/prod/soracom-api.en.yaml"
fja="$d/../soracom-api/dist/prod/soracom-api.ja.yaml"
sen="$d/../soracom-api/dist/sandbox/soracom-sandbox-api.en.yaml"
sja="$d/../soracom-api/dist/sandbox/soracom-sandbox-api.ja.yaml"

if [ ! -f "$fen" ] || [ ! -f "$fja" ] || [ ! -f "$sen" ] || [ ! -f "$sja" ]; then
  echo "API definition files not found. Build soracom-api first."
  exit 1
fi

set -x
npx api-spec-converter --from openapi_3 --to swagger_2 --syntax yaml "$fen" > "$d/generators/assets/soracom-api.en.yaml"
npx api-spec-converter --from openapi_3 --to swagger_2 --syntax yaml "$fja" > "$d/generators/assets/soracom-api.ja.yaml"
npx api-spec-converter --from openapi_3 --to swagger_2 --syntax yaml "$sen" > "$d/generators/assets/sandbox/soracom-sandbox-api.en.yaml"
npx api-spec-converter --from openapi_3 --to swagger_2 --syntax yaml "$sja" > "$d/generators/assets/sandbox/soracom-sandbox-api.ja.yaml"
set +x

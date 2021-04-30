#!/usr/bin/env bash
VERSION=$1
if [ -z "$VERSION" ]; then
  echo "Version number (e.g. 1.2.3) must be specified. Abort."
  usage
  exit 1
fi

AWS_PROFILE=$2
if [ -z "$AWS_PROFILE" ]; then
  AWS_PROFILE=soracom-dev  # 'soracom-dev' is for testing, as a safe default. if you want to go production, specify 'registry'
fi

set -Eeuo pipefail
d="$( cd "$( dirname "$0" )" && cd .. && pwd -P )"

usage() {
  echo "Usage: $0 <version> [aws profile name]"
}

publish_layer() {
  layer_name=$1
  region=$2

  cd "$d" && \
  aws lambda publish-layer-version \
    --layer-name "$layer_name" \
    --zip-file "fileb://soracom/dist/$VERSION/lambda-layer/layer_${VERSION}.zip" \
    --profile "$AWS_PROFILE" \
    --region "$region" \
    --no-cli-pager \
  && cd -

  layer_version="$(
  aws lambda list-layer-versions \
    --layer-name "$layer_name" \
    --profile "$AWS_PROFILE" \
    --region "$region" | jq -r '.LayerVersions[].Version' | sort -n -r | head -n 1
  )"

  aws lambda add-layer-version-permission \
    --layer-name "$layer_name" \
    --version-number "$layer_version" \
    --statement-id AllowGetLayerVersion\
    --principal '*' \
    --action lambda:GetLayerVersion \
    --profile "$AWS_PROFILE" \
    --region "$region" \
    --no-cli-pager
}

should_skip() {
  region=$1

  if [ "$region" == "ap-east-1" ]; then
    return 0
  fi

  return 1
}

layer_name="soracom-cli-${VERSION//./}"
while read -r region; do
  if should_skip "$region"; then
    continue
  fi

  publish_layer "$layer_name" "$region"
done <<< "$( aws ec2 describe-regions --profile "$AWS_PROFILE" --region ap-northeast-1 | jq -r '.Regions[].RegionName' )"

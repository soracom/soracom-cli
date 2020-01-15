#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )"; cd ..; pwd )"
set -e

usage() {
  echo "Usage: $0 <version> <aws profile name>"
}

VERSION=$1
if [ -z "$VERSION" ]; then
  echo "Version number (e.g. 1.2.3) must be specified. Abort."
  usage
  exit 1
fi

AWS_PROFILE=$2
if [ -z "$AWS_PROFILE" ]; then
  echo "AWS profile name must be specified. Abort."
  usage
  exit 1
fi



publish_layer() {
  layer_name=$1
  region=$2

  cd "$d" && \
  aws lambda publish-layer-version \
    --layer-name "$layer_name" \
    --zip-file "fileb://soracom/dist/$VERSION/lambda-layer/layer_${VERSION}.zip" \
    --profile "$AWS_PROFILE" \
    --region "$region" \
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
    --statement-id publishToTheWorld \
    --principal '*' \
    --action lambda:GetLayerVersion \
    --profile "$AWS_PROFILE" \
    --region "$region"

}


layer_name="soracom-cli-${VERSION//./}"
while read -r region; do
  if [ "$region" != "ap-northeast-1" ]; then
    echo "skipping region: $region"
    continue
  fi

  publish_layer "$layer_name" "$region"
done <<< "$( aws ec2 describe-regions --profile "$AWS_PROFILE" --region ap-northeast-1 | jq -r '.Regions[].RegionName' )"

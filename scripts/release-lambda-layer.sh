#!/usr/bin/env bash
set -Eeuo pipefail
d="$( cd "$( dirname "$0" )" && cd .. && pwd -P )"

usage() {
  echo "Usage: $0 <version> [aws profile name]"
}

VERSION=${1:-}
if [ -z "$VERSION" ]; then
  echo "Version number (e.g. 1.2.3) must be specified. Abort."
  usage
  exit 1
fi

AWS_PROFILE=${2:-}
if [ -z "$AWS_PROFILE" ]; then
  AWS_PROFILE=soracom-dev  # 'soracom-dev' is for testing, as a safe default. if you want to go production, specify 'registry'
fi

publish_layer() {
  local layer_name=$1
  local region=$2
  local arch=$3

  local compatible_arch=$arch
  if [ "$compatible_arch" == "amd64" ]; then
    compatible_arch="x86_64"
  fi

  cd "$d" && \
  aws lambda publish-layer-version \
    --layer-name "$layer_name" \
    --zip-file "fileb://soracom/dist/$VERSION/lambda-layer/layer_${VERSION}_${arch}.zip" \
    --description "soracom-cli version $VERSION" \
    --compatible-architectures "$compatible_arch" \
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

while read -r region; do
  if should_skip "$region"; then
    continue
  fi

  layer_name_amd64="soracom-cli-${VERSION//./}"
  publish_layer "$layer_name_amd64" "$region" "amd64"

  layer_name_arm64="soracom-cli-${VERSION//./}-arm64"
  publish_layer "$layer_name_arm64" "$region" "arm64"
done <<< "$( aws ec2 describe-regions --profile "$AWS_PROFILE" --region ap-northeast-1 | jq -r '.Regions[].RegionName' )"

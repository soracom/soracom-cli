#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )"; cd ..; pwd )"
set -e

usage() {
  echo "Usage: $0 <version> <aws profile name> <aws region>"
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

AWS_REGION=$3
if [ -z "$AWS_REGION" ]; then
  echo "AWS region name must be specified. Abort."
  usage
  exit 1
fi


cd "$d" && \
aws lambda publish-layer-version \
  --layer-name "soracom-cli-${VERSION//./}" \
  --zip-file "fileb://soracom/dist/$VERSION/lambda-layer/layer_${VERSION}.zip" \
  --profile "$AWS_PROFILE" \
  --region "$AWS_REGION"

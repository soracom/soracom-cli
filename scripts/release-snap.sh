#!/usr/bin/env bash
VERSION=$1
if [ -z "$1" ]; then
  echo "Version number (e.g. 1.2.3) must be specified. Abort."
  exit 1
fi

set -Eeuo pipefail
d="$( cd "$( dirname "$0" )" && cd .. && pwd )"

docker run -it --rm -v="$d/soracom/dist/$VERSION:/mnt" snapbuild sh -c "bash /opt/release.sh $VERSION"

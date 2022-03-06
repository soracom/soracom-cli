#!/usr/bin/env bash
VERSION=$1
if [ -z "$1" ]; then
  echo "Please specify version number (e.g. 1.2.3)"
  exit 1
fi

set -Eeuo pipefail
d=$( cd "$( dirname "$0" )" && cd .. && pwd -P )

docker build -t snapbuild "$d/snap"
docker run -it --rm -v="$d/soracom/dist/$VERSION:/mnt" snapbuild sh -c "bash /opt/build.sh $VERSION"

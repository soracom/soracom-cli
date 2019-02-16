#!/usr/bin/env bash
d=$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd -P )

VERSION=$1
if [ -z "$1" ]; then
  echo "Please specify version number (e.g. 1.2.3)"
  exit 1
fi

: "Generate snapcraft.yaml" && {
  mkdir -p "$d/soracom/dist/$VERSION/snap"
  cat <<EOD > "$d/soracom/dist/$VERSION/snap/snapcraft.yaml"
name: soracom
version: '$VERSION'
summary: A CLI tool to use SORACOM API.
description: |
  This is a tool to invoke SORACOM API. You can control SORACOM platform
  and its services automatically by using this tool.

grade: stable
confinement: strict

parts:
  soracom:
    plugin: dump
    source: soracom_${VERSION}_linux_amd64.tar.gz
    
apps:
  soracom:
    command: soracom
EOD
}

docker build -t snapbuild "$d/snap"
docker run -it --rm -v="$d/soracom/dist/$VERSION:/mnt" snapbuild sh -c 'bash /opt/build.sh'

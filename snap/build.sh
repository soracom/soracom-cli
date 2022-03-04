#!/usr/bin/env bash
set -Eeuo pipefail
set -x

VERSION=$1

cd /mnt
ls

generate_snapcraft_yaml() {
  local arch=$1
  goarch=$arch
  if [ "$goarch" == "armhf" ]; then
    goarch="arm"
  fi

  mkdir -p "/mnt/snap"
  cat <<EOD > "/mnt/snap/snapcraft.yaml"
name: soracom
version: '$VERSION'
summary: A CLI tool to use SORACOM API.
description: |
  This is a tool to invoke SORACOM API. You can control SORACOM platform
  and its services automatically by using this tool.

grade: stable
confinement: strict
architectures:
  - build-on: all
    run-on: ${arch}

parts:
  soracom:
    plugin: dump
    source: soracom_${VERSION}_linux_${goarch}.tar.gz
    override-build: |
      chmod +x soracom
      snapcraftctl build

apps:
  soracom:
    command: soracom
EOD
}

for arch in amd64 arm64 armhf; do
  generate_snapcraft_yaml "$arch"
  snapcraft clean soracom -s pull
  snapcraft
done

snapcraft clean

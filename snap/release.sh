#!/usr/bin/env bash
set -Eeuo pipefail
set -x

VERSION=$1

snapcraft login

for arch in amd64 arm64 armhf; do
  res="$( snapcraft upload "/mnt/soracom_${VERSION}_${arch}.snap" )"
  rev="$( echo "$res" | sed -rn 's/.*Revision ([0-9]+) of .* created./\1/p' )"
  channels=beta,edge,candidate,stable
  snapcraft release soracom "$rev" "$channels"
done

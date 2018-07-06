#!/bin/bash

VERSION=$1

snapcraft login

set -e
set -x

res="$( snapcraft push "/mnt/soracom_${VERSION}_amd64.snap" )"
rev="$( echo "$res" | sed -rn 's/.*Revision ([0-9]+) of .* created./\1/p' )"
channels=beta,edge,candidate,stable
snapcraft release soracom "$rev" "$channels"

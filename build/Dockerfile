# The Go version used for the build is the one specified by the `go` directive
# in go.mod (scripts/build.sh passes it as GOTOOLCHAIN). Keep the version of this
# image in sync with it; otherwise the toolchain is downloaded on every build.
#
# The base image is pinned by digest so that the release build does not change
# even if the tag gets re-pointed. Update the tag and the digest together:
#   docker buildx imagetools inspect public.ecr.aws/docker/library/golang:<version>
FROM public.ecr.aws/docker/library/golang:1.27.4@sha256:5e8afd493a7124a8586e134d9b532731161d4b9becf9dee5a5956b107426f93d

RUN apt update && apt install -y fakeroot shellcheck zip

ADD build.sh /build/

CMD [ "bash", "-x", "/build/build.sh" ]

#!/usr/bin/env bash
d=$( cd "$( dirname "$0" )"; cd ..; pwd -P )

: 'Check if shell scripts are healthy' && {
  command -v shellcheck > /dev/null 2>&1 && {
    shellcheck -e SC2164 "$d/scripts/"*.sh
    shellcheck -e SC2164 "$d/test/"*.sh
  }
}

goversion=1.15
docker pull "golang:$goversion"
gopath=${GOPATH:-$HOME/go}
gopath=${gopath%%:*}

run_command_on_docker_container() {
  dir=$1
  cmd=$2
  #echo $cmd
  if [ -z "$WERCKER" ]; then
    docker run -i --rm \
      --user "$(id -u):$(id -g)" \
      -v "$d":/go/src/github.com/soracom/soracom-cli \
      -v "$gopath":/go \
      -v "$d/.cache":/.cache \
      -w "/go/src/github.com/soracom/soracom-cli/$dir" \
      "golang:$goversion" bash -x -c "$cmd" || {
      echo -e "${RED}Build failed.${RESET}"
      exit 1
    }
  else
    # on wercker, it's already running on a docker container
    set -x
    cd "/go/src/github.com/soracom/soracom-cli/$dir" && GO111MODULE=on bash -c "$cmd"
    set +x
  fi
}

set -e # aborting if any commands below exit with non-zero code

VERSION=$1
if [ -z "$1" ]; then
  VERSION='0.0.0'
  echo "Version number (e.g. 1.2.3) is not specified. Using $VERSION as the default version number"
fi

TARGETS=$2
if [ -z "$2" ]; then
    TARGETS='linux windows darwin,!386 freebsd'
    uname_s="$( uname -s | tr '[:upper:]' '[:lower:]' )"
    if [[ "$TARGETS" != *"$uname_s"* ]]; then
        TARGETS="$TARGETS $uname_s"
    fi
fi

: 'Install dependencies' && {
    echo 'Installing build dependencies ...'
    run_command_on_docker_container '' 'go get -u golang.org/x/tools/cmd/goimports'
    run_command_on_docker_container '' 'go get -u github.com/laher/goxc'

    echo 'Installing commands used with "go generate" ...'
    run_command_on_docker_container '' 'go get -u github.com/jessevdk/go-assets-builder'
    run_command_on_docker_container '' 'go get -u github.com/elazarl/goproxy'
    run_command_on_docker_container '' 'go mod tidy'
}

: "Test generator's library" && {
    echo "Testing generator's source ..."
    run_command_on_docker_container 'generators/cmd/src' 'go test'
    run_command_on_docker_container 'generators/lib'     'go test'
}

: 'Generate source code for soracom-cli' && {
    echo 'Generating generator ...'
    run_command_on_docker_container 'generators/cmd/src' 'go generate'
    run_command_on_docker_container 'generators/cmd/src' 'go vet'
    run_command_on_docker_container 'generators/cmd/src' 'goimports -w ./*.go'
    run_command_on_docker_container 'generators/cmd/src' 'go test'
    run_command_on_docker_container 'generators/cmd/src' 'go build -o generate-cmd'

    echo 'Generating source codes for soracom-cli by using the generator ...'
    run_command_on_docker_container '' 'generators/cmd/src/generate-cmd -a generators/assets/soracom-api.en.yaml -s generators/assets/sandbox/soracom-sandbox-api.en.yaml -t generators/cmd/templates -p generators/cmd/predefined -o soracom/generated/cmd/'
}

: 'Build soracom-cli executables' && {
    echo 'Building artifacts ...'
    run_command_on_docker_container 'soracom' 'go generate'
    run_command_on_docker_container 'soracom' 'go get -u github.com/bearmini/go-acl' # required to specify some dependencies explicitly as they are imported only in windows builds
    run_command_on_docker_container 'soracom' 'gofmt -s -w .'
    run_command_on_docker_container 'soracom' "goxc -bc='$TARGETS' -d=dist/ -pv=$VERSION -build-ldflags='-X github.com/soracom/soracom-cli/soracom/generated/cmd.version=$VERSION' -tasks-=rmbin"

    # non-zipped versions for homebrew and testing
    echo 'Renaming artifacts for homebrew ...'
    for distfile in "$d/soracom/dist/$VERSION"/*/soracom; do
      distos="${distfile%/*}"
      distos="${distos##*/}"
      mv "$distfile" "$d/soracom/dist/$VERSION/soracom_${VERSION}_${distos}"
    done
    for distfile in "$d"/soracom/dist/"$VERSION"/*/soracom.exe; do
      distos="${distfile%/*}"
      distos="${distos##*/}"
      mv "$distfile" "$d/soracom/dist/$VERSION/soracom_${VERSION}_${distos}.exe"
    done

    # removing empty directories
    find "$d/soracom/dist/$VERSION" -type d -empty -delete
}

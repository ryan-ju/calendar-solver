#!/usr/bin/env bash

# Exit on first error
set -e

export CUR_DIR="$(pwd)"

print_help() {
    echo "usage: build.sh fix|lint|test|docker-build|help"
    echo "  fix           auto format Go source code and tidy go.sum for each sub module in this project"
    echo "  lint          lint each sub module in this project"
    echo "  test          run test in each sub module"
    echo "  docker-build  build docker image for api-cli"
    echo "  help          print this message"
    echo ""
    echo "Make sure golangci-lint is installed, by running:"
    echo "    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint"
    exit 1
}

list_dirs() {
    find . -name '*.go' -print0 | xargs -0 -n1 dirname | sort --unique | grep -v vendor | grep -v tmp | grep -v "_static"
}

fix() {
    echo "Fixing imports ..."
    goimports -l -w $(list_dirs)
    echo "Tidying go.sum ..."
    for d in $(list_dirs); do
        echo "Tidying directory $d ..."
        cd "$d"
        go mod tidy
        cd "$CUR_DIR"
    done
}

lint() {
    golangci-lint run ./...
}

go_test() {
    go test ./...
}

# For CircleCI only
circle_test() {
    go test -tags circle ./...
}

install() {
    go install ./cmd/api-cli
}

build_docker() {
    mkdir -p dist
    GOOS=linux GOARCH=amd64 go build -o dist/api-cli-core ./cmd/api-cli-core
    GOOS=linux GOARCH=amd64 go build -o dist/protoc-gen-trinny-gateway ./cmd/protoc-gen-trinny-gateway
    docker build -t trinnylondon/api-cli .
}

case "$1" in
    help)
        print_help
    ;;
    fix)
        fix
    ;;
    lint)
        lint
    ;;
    test)
        go_test
    ;;
    circle-test)
        circle_test
    ;;
    install)
        install
    ;;
    build-docker)
        build_docker
    ;;
    *)
        print_help
    ;;
esac

exit 0
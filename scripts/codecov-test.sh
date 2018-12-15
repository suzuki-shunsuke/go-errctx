#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

go test -race -coverprofile=coverage.out -covermode=atomic . || exit 1
test "$LOCAL" == "true" || bash <(curl -s https://codecov.io/bash) || exit 1

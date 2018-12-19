#!/bin/sh

GOFMT=$(gofmt -l .)
if [ -n "${GOFMT}" ]; then
    printf >&2 'gofmt failed for:\n%s\n' "${GOFMT}"
    exit 1
fi

GOTEST=$(go test ./...)
if [ $? -ne 0 ]; then
    printf >&2 'go test failed:\n%s\n' "${GOTEST}"
    exit 1
fi

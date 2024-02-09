#!/bin/sh

set -e
COMMAND=$@

echo 'Testing...'
go test -v ./...


exec $COMMAND


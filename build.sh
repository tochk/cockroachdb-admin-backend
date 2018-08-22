#!/bin/bash

set -xe
set -u

mkdir -p build

env GOOS=linux GOARCH=amd64 go build -o build/cdbadmin-linux-amd64
env GOOS=windows GOARCH=amd64 go build -o build/cdbadmin-windows-amd64.exe
env GOOS=darwin GOARCH=amd64 go build -o build/cdbadmin-darwin-amd64
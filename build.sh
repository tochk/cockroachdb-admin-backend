#!/bin/bash

set -xe
set -u

mkdir -p build

gox -arch="386 amd64" -os="darwin windows linux" -output="./build/crdbadmin_{{.OS}}_{{.Arch}}"
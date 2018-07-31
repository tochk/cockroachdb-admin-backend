#!/usr/bin/env bash

set -xe
set -u

rm -rf static/index.html
rm -rf static/build.js
rm -rf static/build.js.map
rm -rf tmp

mkdir tmp
cd tmp

git clone https://github.com/tochk/cockroachdb-admin.git
cd cockroachdb-admin/vue-spa

yarn install
yarn build

mv index.html ../../../static/index.html
mv dist/build.js ../../../static/build.js
mv dist/build.js.map ../../../static/build.js.map

cd ../../../
rm -rf tmp

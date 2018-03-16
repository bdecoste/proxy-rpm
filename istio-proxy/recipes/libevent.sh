#!/bin/bash

set -e

VERSION=2.1.8-stable

pwd 
ls -l
echo "${RPM_BUILD_DIR}"

#wget -O libevent-"$VERSION".tar.gz https://github.com/libevent/libevent/releases/download/release-"$VERSION"/libevent-"$VERSION".tar.gz
#tar xf libevent-"$VERSION".tar.gz

cp -rf ${RPM_BUILD_DIR}/istio-proxy/libevent-"$VERSION" .

cd libevent-"$VERSION"
./configure --prefix="$THIRDPARTY_BUILD" --enable-shared=no --disable-libevent-regress --disable-openssl
make V=1 install

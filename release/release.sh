#!/bin/bash

# Copyright (c) 2016 Company 0, LLC.
# Copyright (c) 2016-2020 The btcsuite developers
# Use of this source code is governed by an ISC
# license that can be found in the LICENSE file.

# Simple bash script to build basic btcd tools for all the platforms we support
# with the golang cross-compiler.

set -e

export TZ=UTC
umask 022

# If no tag specified, use date + version otherwise use tag.
if [[ $1x = x ]]; then
    DATE=`date +%Y%m%d`
    VERSION="01"
    TAG=$DATE-$VERSION
else
    TAG=$1
fi

TAR_OPTS="-H pax --sort=name --mtime=1970-01-01 --owner=0 --group=0 --numeric-owner --pax-option=exthdr.name=%d/PaxHeaders/%f,delete=atime,delete=ctime"
go mod vendor
tar $TAR_OPTS -cvf vendor.tar vendor
gzip -n6 vendor.tar
shasum -a 256 vendor.tar.gz

PACKAGE=grsd
MAINDIR=$PACKAGE-$TAG
mkdir -p $MAINDIR

cp vendor.tar.gz $MAINDIR/
rm vendor.tar.gz
rm -r vendor

PACKAGESRC="$MAINDIR/$PACKAGE-source-$TAG.tar"
git archive -o $PACKAGESRC HEAD
gzip -fn6 $PACKAGESRC > $PACKAGESRC.gz
shasum -a 256 $PACKAGESRC.gz

cd $MAINDIR

# If GRSDBUILDSYS is set the default list is ignored. Useful to release
# for a subset of systems/architectures.
SYS=${GRSDBUILDSYS:-"
        darwin-amd64
        dragonfly-amd64
        freebsd-386
        freebsd-amd64
        freebsd-arm
        illumos-amd64
        linux-386
        linux-amd64
        linux-armv6
        linux-armv7
        linux-arm64
        linux-ppc64
        linux-ppc64le
        linux-mips
        linux-mipsle
        linux-mips64
        linux-mips64le
        linux-s390x
        netbsd-386
        netbsd-amd64
        netbsd-arm
        netbsd-arm64
        openbsd-386
        openbsd-amd64
        openbsd-arm
        openbsd-arm64
        solaris-amd64
        windows-386
        windows-amd64
"}
#SYS="linux-amd64 windows-amd64"

for i in $SYS; do
    OS=$(echo $i | cut -f1 -d-)
    ARCH=$(echo $i | cut -f2 -d-)
    ARM=

    if [[ $ARCH = "armv6" ]]; then
      ARCH=arm
      ARM=6
    elif [[ $ARCH = "armv7" ]]; then
      ARCH=arm
      ARM=7
    fi

    mkdir $PACKAGE-$i-$TAG
    cd $PACKAGE-$i-$TAG

    if [[ $OS = "windows" ]]; then
        GRSD=grsd.exe
        GRSCTL=grsctl.exe
    else
        GRSD=grsd
        GRSCTL=grsctl
    fi
    echo "Building:" $OS $ARCH $ARM
    env CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH GOARM=$ARM go build -v -trimpath -ldflags="-s -w -buildid=" -o $GRSD github.com/btcsuite/btcd
    env CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH GOARM=$ARM go build -v -trimpath -ldflags="-s -w -buildid=" -o $GRSCTL github.com/btcsuite/btcd/cmd/btcctl
    cd ..

    if [[ $OS = "windows" ]]; then
        zip -X -r $PACKAGE-$i-$TAG.zip $PACKAGE-$i-$TAG
        strip-nondeterminism $PACKAGE-$i-$TAG.zip
    else
        tar $TAR_OPTS -cvf $PACKAGE-$i-$TAG.tar $PACKAGE-$i-$TAG
        gzip -fn6 $PACKAGE-$i-$TAG.tar
    fi

    rm -r $PACKAGE-$i-$TAG
    shasum -a 256 $PACKAGE-$i-$TAG.*
done

shasum -a 256 * > manifest-$TAG.txt
cat manifest-$TAG.txt

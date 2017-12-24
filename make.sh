#!/usr/bin/env bash


basedir=$(cd -- "${BASH_SOURCE%/*}" && pwd)

rm -rf -- "$basedir/"{build}
mkdir -p -- "$basedir/"{build}

build() (  GOOS=$1 GOARCH=$2 exec go build -o "$basedir/build/$3")

build darwin  amd64 mzr_mac
build linux   amd64 mzr
build windows amd64 mzr_win64


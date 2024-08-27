#!/bin/bash

set -x -e

download() {
  curl -L "https://github.com/${RELEASE_REPO}/releases/download/$1/$2" >"$2"
}

downloadAndUnzip() {
  download "$1" "$2"
  unzip -n -d "${2%\.zip}" "$2"
}
mkdir -p assets

pushd assets
downloadAndUnzip "$1" "accelerator-linux-$2.zip"
downloadAndUnzip "$1" "accelerator-extra.zip"
popd

placeFile() {
  mkdir -p "context/$2"
  cp -R "assets/$1/$3" "context/$2/$3"
}

function generateStandardVersion() {
  placeFile "$1" "$2/bin" "accelerator"
}

function generateExtraVersion() {
  generateStandardVersion "$1" "$2"
  placeFile "$1" "$2/share" "geosite.dat"
  placeFile "$1" "$2/share" "geoip.dat"
  placeFile "$1" "$2/etc" "config.json"
  placeFile "accelerator-extra" "$2/share" "browserforwarder"
}

if [ "$4" = "std" ]; then
    generateStandardVersion "accelerator-linux-$2" "linux/$3/std"
fi

if [ "$4" = "extra" ]; then
    generateExtraVersion "accelerator-linux-$2" "linux/$3/extra"
fi



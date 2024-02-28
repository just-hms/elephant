#!/bin/bash

if [ -z "$1" ]; then
    echo "Error: Version not specified."
    exit 1
fi

VERSION=$1
BUILD="./build/"
OUT="./out/"

echo "Creating the build folders..."
mkdir -p $OUT
mkdir -p $BUILD

echo "Copying shell files to build folder..."
cp -r ./shell/. $BUILD

echo "Building..."

echo "linux"
env GOARCH=amd64 GOOS=linux go build -ldflags="-X 'el/internal/constants.Version=$VERSION'" -o $BUILD/el ./cmd/
tar -cvzf $OUT/el-linux-amd64.tar.gz -C $BUILD .

echo "darwin"
env GOARCH=amd64 GOOS=darwin go build -ldflags="-X 'el/internal/constants.Version=$VERSION'" -o $BUILD/el ./cmd/
tar -cvzf $OUT/el-darwin-amd64.tar.gz -C $BUILD .

# echo "windows"
# env GOARCH=amd64 GOOS=windows go build -ldflags="-X 'el/internal/constants.Version=$1'" -o $BUILD/el ./cmd/
# tar -cvzf $OUT/el-darwin-amd64.tar.gz -C $BUILD .

echo "Cleaning up..."
rm -r $BUILD

echo "Outputs:"
ls -l $OUT

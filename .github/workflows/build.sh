#!/bin/bash

if [ -z "$1" ]; then
    echo "Error: Version not specified."
    exit 1
fi

VERSION=$1
BUILD="./build"
OUT="./out"

echo "Creating the build folders..."
mkdir -p $OUT
mkdir -p $BUILD

echo "Copying shell files to build folder..."
cp ./shell/* $BUILD

echo "Building..."

echo "linux"
env GOARCH=amd64 GOOS=linux go build -ldflags="-X 'el/internal/constants.Version=$VERSION'" -o $BUILD/el ./cmd/
tar -cvzf $OUT/el-linux-amd64.tar.gz $BUILD/*

echo "darwin"
env GOARCH=amd64 GOOS=darwin go build -ldflags="-X 'el/internal/constants.Version=$VERSION'" -o $BUILD/el ./cmd/
tar -cvzf $OUT/el-darwin-amd64.tar.gz $BUILD/*

# echo "windows"
# env GOARCH=amd64 GOOS=windows go build -ldflags="-X 'el/internal/constants.Version=$1'" -o ./out/el.exe ./cmd/
# (cd $OUT && tar -cvzf * && rm ./el.exe)

echo "Outputs:"
ls -l $OUT

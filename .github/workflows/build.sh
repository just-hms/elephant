#!/bin/bash

if [ -z "$1" ]; then
    echo "Error: Version not specified."
    exit 1

fi

VERSION=$1
OUT="./out"

echo "Creating the ouput folder..."
mkdir -p $OUT

echo "Copying README.md to output folder..."
cp ./README.md $OUT

echo "Building..."

echo "linux"
env GOARCH=amd64 GOOS=linux go build -ldflags="-X 'el/internal/constants.Version=$VERSION'" -o $OUT/el ./cmd/
(cd $OUT && tar -cvzf el-linux-amd64.tar.gz README.md el && rm ./el)

echo "darwin"
env GOARCH=amd64 GOOS=darwin go build -ldflags="-X 'el/internal/constants.Version=$VERSION'" -o $OUT/el ./cmd/
(cd $OUT && tar -cvzf el-darwin-amd64.tar.gz README.md el && rm ./el)

echo "windows"
env GOARCH=amd64 GOOS=windows go build -ldflags="-X 'el/internal/constants.Version=$1'" -o ./out/el.exe ./cmd/
(cd $OUT && tar -cvzf el-windows-amd64.tar.gz README.md el && rm ./el)


echo "Cleaning up.."
rm $OUT/README.md

echo "Outputs:"
ls -l $OUT

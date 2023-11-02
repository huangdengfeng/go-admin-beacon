#!/bin/bash
set -e
DIR=$(cd "$(dirname $0)" && pwd)
APP_NAME="${DIR##*/}"
cd $DIR
rm -rf target
mkdir -p target/$APP_NAME
\cp -R bin target/$APP_NAME
\cp -R conf target/$APP_NAME
# -o 后加目录则放入目录，不是目录则为产出物名称
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/$APP_NAME/bin/$APP_NAME cmd/*
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd -o build/package/$APP_NAME/bin/$APP_NAME cmd/*
chmod -R +x target/$APP_NAME/bin/
cd target && tar -zcvf $APP_NAME.tar.gz $APP_NAME
#!/bin/bash

cd helperTool
go build -o bin/helperTool
bin/helperTool updateBindata

if [ $? -ne 0 ]; then
  echo "Bindata failed! bailing.."
  exit 1
fi

env GOOS="linux" GOARCH="amd64" go build -o bin/linux/bt-bot
upx bin/linux/bt-bot


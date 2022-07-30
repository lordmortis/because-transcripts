#!/bin/sh

cd helperTool
go build -o bin/helperTool
bin/helperTool updateBindata

env GOOS="linux" GOARCH="amd64" go build -o bin/linux/bt-bot
upx bin/linux/bt-bot


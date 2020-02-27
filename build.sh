#!/bin/bash

# 构建可执行二进制文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goinspect ./server.go

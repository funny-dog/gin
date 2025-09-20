#!/bin/bash
# Setup Go proxy for Docker containers to fix dependency download issues

# China Go proxy (faster for Chinese users)
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

# Alternative proxies if the main one fails
# export GOPROXY=https://goproxy.io,direct
# export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

echo "Go proxy configured to use: $GOPROXY"
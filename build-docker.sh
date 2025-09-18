#!/bin/bash

# 构建 Docker 镜像
echo "开始构建 Docker 镜像..."
docker build -t gin-framework .

if [ $? -eq 0 ]; then
    echo "Docker 镜像构建成功!"
else
    echo "Docker 镜像构建失败!"
    exit 1
fi

# 启动容器
echo "启动容器..."
docker run -d -p 8080:8080 --name gin-app gin-framework

if [ $? -eq 0 ]; then
    echo "容器启动成功!"
    echo "访问以下端点测试:"
    echo "  http://localhost:8080/health"
    echo "  http://localhost:8080/ping"
    echo "  http://localhost:8080/gin/info"
else
    echo "容器启动失败!"
    exit 1
fi
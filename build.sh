#!/bin/bash

# 设置构建的输出目录
BUILD_DIR="build"

# 创建输出目录
mkdir -p $BUILD_DIR

# 执行打包操作
go build -o $BUILD_DIR/your_executable ./server

# 进入输出目录
cd $BUILD_DIR || exit

# 打包成压缩文件
tar -czvf your_project_v1.0.0.tar.gz your_executable

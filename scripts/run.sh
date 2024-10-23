#!/bin/bash

# 输出开始信息
echo "Starting Minty..."

# 切换到 cmd 目录
echo "Changing directory to ../cmd..."
cd ../cmd || { echo "Failed to change directory to ../cmd"; exit 1; }

# 构建应用程序
echo "Building application..."
go build -o ../build/minty main.go || { echo "Failed to build application"; exit 1; }

# 运行应用程序
echo "Running application..."
../build/minty || { echo "Failed to run application"; exit 1; }

# 输出完成信息
echo "Minty started successfully."
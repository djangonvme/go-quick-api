#!/bin/bash

# go env -w GOPRIVATE=gitlab.gosccap.cn


projectRoot=`pwd`

# 二进制名称
binaryName=goquickapi

# 运行的日志
log="${binaryName}.run.log"

# 启动的参数，配置文件
startArg="-config=${projectRoot}/config.ini"

case $1 in
    start)
        pkill  $binaryName
        echo "开始编译"
        make build
        nohup ./bin/$binaryName  "$startArg"  0</dev/null >>$log 2>&1 &
        echo "服务已启动..."
    ;;
    stop)
        pkill  $binaryName
        echo "服务已停止..."
    ;;
    restart)
        pkill $binaryName
        echo "开始编译"
        make build
        nohup ./bin/$binaryName "$startArg" 0</dev/null >>$log 2>&1 &
        echo "服务已重启..."
    ;;
    *)
        echo "$0 {start|stop|restart}"
        exit 4
    ;;
esac
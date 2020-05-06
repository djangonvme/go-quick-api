#!/bin/bash

command=$1
commandArg=$2

configLinkFile="/etc/ginapicommon_config.ini"


if [ ! -n "$action" ];then
    action="up"
fi
usage(){
    echo "Usage:  ./start.sh [action]"
    echo "actions:"
    echo "  up               docker-compose up"
    echo "  build-up         docker-compose 重新构建代码镜像和启动"
    echo "  down             docker-compose down 停止服务"
    echo "  remove-all       移除docker-compose启动相关的容器"
    echo "  local            本地启动，使用本地的配置"
    echo "  fresh            本地启动，实时检测代码变动和重新编译, 适合本地开发debug"
}


# 代码改动后重新构建镜像, 仅仅代码的程序镜像重新构建, 其他服务镜像不动
rebuildUp(){
   # 重新构建镜像
   docker build -t gin-api-common_main:latest -f deploy/docker-compose/ginapicommon.build .
   # 停止旧容器
   docker stop ginapicommon_main  &&  docker rm ginapicommon_main
   docker-compose down
   docker-compose up $commandArg
}
# 移除所有docker-compose启动相关的容器
removeAll(){
    docker-compose down
    docker stop ginapicommon_main ginapicommon_redis  ginapicommon_mysql
    docker rm ginapicommon_main ginapicommon_redis  ginapicommon_mysql
    docker rmi gin-api-common_main gin-api-common_mysql
    docker volume rm gin-api-common_main
    docker volume rm gin-api-common_mysql
}

# 本地启动，使用本地配置
localStart(){
  if [ ! -f "$configLinkFile" ];then
    sudo ln -s $(pwd)/config.ini /etc/ginapicommon_config.ini
  fi
  go run main.go -config /etc/ginapicommon_config.ini
}

# 本地运行， 实时检测代码变动和重新编译
localFreshStart(){
  if [ ! -f "$configLinkFile" ];then
    sudo ln -s $(pwd)/config.ini /etc/ginapicommon_config.ini
  fi
  ./fresh
}

#
case $command in
    up)
        docker-compose up $commandArg
    ;;
    down)
        docker-compose down
    ;;
    build-up)
        rebuildUp
    ;;
    local)
        localStart
    ;;
    fresh)
        localFreshStart
    ;;
    remove-all)
        removeAll
    ;;
    -h)
        usage
    ;;
    *)
        echo "invalid command, please view usage: ./start.sh -h"
        usage
    ;;
esac

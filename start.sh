#!/bin/bash

action=$1

usage(){
    echo "Usage:  ./start.sh [action]"
    echo "actions:"
    echo "  up               equal docker-compose up"
    echo "  down             docker-compose down"
    echo "  restart          rebuild main image and remove main containter then  docker-compose restart"
    echo "  remove-all       remove all!"
}

# 代码改动后重新构建镜像
buildRestart(){
   # 重新构建镜像
   docker build -t gin-api-common_main:latest -f deploy/docker-compose/ginapicommon.build .
   # 停止旧容器
   docker stop ginapicommon_main  &&  docker rm ginapicommon_main
   docker-compose down
   docker-compose up
}


case $action in
    up)
        docker-compose up
    ;;
    down)
        docker-compose down
    ;;
    restart)
        buildRestart
    ;;
    remove-all)
        docker-compose down
        docker stop ginapicommon_main ginapicommon_redis  ginapicommon_mysql
        docker rm ginapicommon_main ginapicommon_redis  ginapicommon_mysql
        docker rmi gin-api-common_main gin-api-common_mysql
    ;;
    -h)
        usage
    ;;
esac



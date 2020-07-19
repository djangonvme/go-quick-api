#!/bin/bash


# image or container name
img_app="goquickapi_main"
img_mysql="goquickapi_mysql"

# stop and remove containers,images
docker stop         "$img_app" "$img_mysql"
docker rm           "$img_app" "$img_mysql"
docker rmi          "$img_app" "$img_mysql"

# build images
docker build -t     "$img_app":latest -f deploy/image/main.build .
docker build -t     "$img_mysql":latest -f deploy/image/mysql.build .

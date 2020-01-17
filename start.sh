#!/bin/bash

docker stop ginapicommon_main ginapicommon_redis  ginapicommon_mysql
docker rm ginapicommon_main ginapicommon_redis  ginapicommon_mysql
docker-compose build --no-cache
docker-compose up

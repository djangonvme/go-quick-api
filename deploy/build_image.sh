#!/bin/bash

docker stop goquickapi_main goquickapi_mysql
docker rm goquickapi_main goquickapi_mysql
docker rmi goquickapi_main goquickapi_mysql




docker build -t goquickapi_main:latest -f deploy/docker-compose/project.build .
docker build -t goquickapi_mysql:latest -f deploy/docker-compose/mysql.build .

#!/bin/bash

docker build -t goquickapi_main:latest -f deploy/docker-compose/project.build .
docker build -t goquickapi_mysql:latest -f deploy/docker-compose/mysql.build .

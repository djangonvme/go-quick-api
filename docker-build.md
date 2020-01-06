Deploy with Docker

> At project root dir, do the operations below by steps:

# 1, build image

```shell script
docker build -t gin-api-common:latest -f ./Dockerfile . 
```

# 2, mysql
```shell script
docker pull mysql
docker run --name mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456  -e TZ=Asia/Shanghai  -d mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci --default-time_zone='+08:00'
```

# 3, redis
```shell script
docker pull redis
docker run -itd --name redis -p 6378:6379 redis
```

## 4, Run App

```shell script
# run container
# mysql-ci is alias for container name mysql, same as redis-ci, then use mysql-ci as connect host in you codes for connectting mysql 
# for example: connectMysql("mysql-ci:3306", "root", "123456"), the port 3306 has no business with mysql container exposed ports
# Actually the alias mysql-ci means a ip in docker container, every container has a ip
# tip: alias mysql-ci mustn't be added to /etc/host as domain redirect to any ip in your original computer!
docker run -itd --name gin-api-common  --link mysql:mysql-ci --link redis:redis-ci -p 8081:8080 gin-api-common:latest
# logs
docker logs gin-api-common

```

## Rebuild

```shell script

# remove and rebuild
docker stop gin-api-common && docker rm gin-api-common && docker rmi gin-api-common:latest && docker build -t gin-api-common:latest -f ./Dockerfile .

docker run -itd --name gin-api-common  --link mysql:mysql-ci --link redis:redis-ci -p 8081:8080 gin-api-common:latest 
# see start logs
docker logs gin-api-common
# entry try one way
docker exec -it gin-api-common sh 
docker exec -it gin-api-common /bin/bash
docker exec -it gin-api-common bash
#--------

```

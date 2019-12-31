Deploy with Docker

> At project root dir, do the operations below by steps:

# 1, build image

```shell script
docker build -t gindemo:v1 -f ./Dockerfile . 
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
docker run -itd --name gindemo  --link mysql:mysql-ci --link redis:redis-ci -p 8081:8080 gindemo:v1
# logs
docker logs gindemo

```

## Rebuild

```shell script

# remove and rebuild
docker stop gindemo && docker rm gindemo && docker rmi gindemo:v1 && docker build -t gindemo:v1 -f ./Dockerfile .

docker run -itd --name gindemo  --link mysql:mysql-ci --link redis:redis-ci -p 8081:8080 gindemo:v1
# see start logs
docker logs gindemo

```

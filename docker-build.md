Deploy with Docker


> At project root dir, do the operations below by steps:


## 1, Build image

```shell script
docker build -t gindemo:v1 -f ./Dockerfile . 
```


## 2, run mysql container
```shell script
docker pull mysql
docker run --name mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```

## 3, run redis container
```shell script
docker pull redis
docker run -itd --name redis -p 6378:6379 redis
```

## 4, Run App container

```shell script
docker run -itd --name gindemo  --link mysql:mysql --link redis:redis -p 8081:8080 gindemo:v1
```


## Rebuild

```shell script

## remove and rebuild
docker stop gindemo && docker rm gindemo && docker rmi gindemo:v1 && docker build -t gindemo:v1 -f ./Dockerfile .
## run container
docker run -itd --name gindemo  --link mysql:mysql --link redis:redis -p 8081:8080 gindemo:v1

```

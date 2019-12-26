Deploy with Docker

> At project root dir, do the operations below by steps:


## 1, Build centos image

```shell script
docker build -t centos-go:v1 -f deploy/centos-go.build . 
```

## 2, build app image
```shell script
docker build -t ginapp:v1 -f deploy/app.build . 

```

## 3, run msyql container
```shell script
docker pull mysql
docker run --name mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```

## 4, reids
```shell script
docker pull redis
docker run -itd --name redis -p 6378:6379 redis
```

## 5, Run App container

```shell script
##设置-v可能有权限问题
##docker run -itd --name ginapp -p 8081:8080 -v /logs:/logs -v ~/www/go-projects:/data/projects  --privileged=true centos-go:v7 /usr/sbin/init
docker run -itd --name ginapp-demo  --link mysql:mysql --link redis:redis -u 501:501 -p 8081:8080  ginapp:v1
```

## 6, start app
```shell script
docker exec ginapp-demo ./start.sh 
```


## Rebuild

```shell script
## remove and rebuild
docker stop ginapp-demo && docker rm ginapp-demo && docker rmi ginapp:v1 && docker build -t ginapp:v1 -f deploy/app.build .
## run container
docker run -itd --name ginapp-demo  --link mysql:mysql --link redis:redis  -p 8081:8080 ginapp:v1
## start app
docker exec ginapp-demo ./start.sh 
```

Deploy your App with Docker

> At project root dir

## Build centos env

```shell script
docker build -t centos-go:v1 -f deploy/centos-go.build .

```

## Build app

```shell script

docker build -t ginapp:v1 -f deploy/app.build .

```


## Run container

```shell script

docker run -itd --name ginapp -p 8081:8080 -v /logs:/logs ginapp

```



## Rebuild ginapp

```shell script
docker stop ginapp
docker rm ginapp
docker rmi ginapp
docker build -t ginapp:v1 -f deploy/app.build .
```









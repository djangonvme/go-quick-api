mysql, redis 等有状态的服务自己修改配置是否容器运行即可，这里举例以容器方式运行

在项目根目录：

1, 构建项目代码镜像

```shell script
docker build -t goquickapi_main:latest -f deploy/docker-compose/project.build .
```
2, 构建mysql镜像(含初始化数据)
```shell script
docker build -t goquickapi_mysql:latest -f deploy/docker-compose/mysql.build . 
```
3, 启动

```shell script
docker-compose up
 ```
或
```shell script
docker-compose up -d
```
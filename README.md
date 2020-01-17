# gin-api-common 

* 基于go web框架gin的简单封装, 因为gin没有项目目录等传统框架定义的东西, 封装的目的是扩展一些现代接口开发的必备功能, 提高开发效率, 你可以拿来即用。

* 1小时入门go web开发， 为go初学者展示如何使用gin框架快速web开发，这个项目有完整示范: 
启动server->路由定义->中间件过滤->用户身份验证->控制器处理->数据库操作->redis缓存->返回给客户端, 这是web接口开发基本流程

# 功能特色
* gin框架为基础。文档： https://github.com/gin-gonic/gin
* 数据库操作orm 用 gorm。 文档： http://gorm.book.jasperxu.com 
* 使用 JWT 生成token, 结合redis双重验证。 文档 http://jwt.io
* gin 使用的验证器文档: https://godoc.org/gopkg.in/go-playground/validator.v8
* 日志用logrus  文档: https://github.com/sirupsen/logrus
* docker 部署, docker-compose 编排容器一键启动整套服务

# 版本要求

 * go1.11+
 * 使用 go modules 包管理

# 项目结构

```cassandraql
├── Dockerfile
├── Makefile
├── README.md
├── apis
│   └── v1
│       ├── exampleLoginApi.go
│       └── exampleUserApi.go
├── config.ini
├── consts
│   ├── api.go
│   ├── common.go
│   └── redis.go
├── go.mod
├── go.sum
├── libs
│   ├── config.go
│   ├── db.go
│   ├── logger.go
│   └── redis.go
├── main.go
├── middlewares
│   ├── api.go
│   ├── common.go
│   └── log.go
├── models
│   ├── model.go
│   ├── user.go
│   └── userToken.go
├── params
│   └── userApi.go
├── routes
│   └── api.go
├── services
│   ├── login.go
│   └── user.go
└── utils
    ├── encrypt.go
    ├── http.go
    ├── jwt.go
    ├── jwt_test.go
    ├── response.go
    ├── time.go
    └── var.go


```

# 一键启动
根目录执行： 

```
docker-compose up
```
启动完成后 docker ps 查看已经启动了3个容器：
```cassandraql
$ docker ps

CONTAINER ID        IMAGE                COMMAND                  CREATED             STATUS              PORTS                               NAMES
7c6e7d32b784        ginapicommon_web     "dockerize -wait tcp…"   24 minutes ago      Up 24 minutes       0.0.0.0:8080->8080/tcp              ginapicommon_main
f32c1cf52315        redis:latest         "docker-entrypoint.s…"   24 minutes ago      Up 24 minutes       0.0.0.0:6380->6379/tcp              ginapicommon_redis
76382ed6fde1        ginapicommon_mysql   "docker-entrypoint.s…"   24 minutes ago      Up 24 minutes       33060/tcp, 0.0.0.0:3309->3306/tcp   ginapicommon_mysql
```



打开浏览器访问： http://127.0.0.1:8080 看效果
```cassandraql
⇒  curl http://127.0.0.1:8080/
{"code":200,"msg":"请求成功","timestamp":1559253308,"data":"Welcome!"}%
```

# 重新全部构建并启动：

```cassandraql
docker stop ginapicommon_main ginapicommon_redis  ginapicommon_mysql
docker rm ginapicommon_main ginapicommon_redis  ginapicommon_mysql
docker-compose build --no-cache
docker-compose up
```


# 请求示例

* 添加用户
路由在/routes/api.go 中可以看到 

 
* 登陆


路由在/routes/api.go 中可以看到 ```/login``` 




参数:
```json
{"mobile": "1500000000", "pwd": "1234546"}
```
响应:
```json
{
    "code": 200,
    "msg": "请求成功",
    "timestamp": 1576573879,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIsInVlbiI6ImVjNDc2ZDJkNGU3ODhkYzA3YzFkNDI3NGVkZjA1Y2Y1YmQyMGI4YWYwYTdlODcwYTAzMzRmYjZlZDg2MzNiZDQiLCJleHAiOjE1NzY2NjAyNzksImlzcyI6InRlc3QifQ.erealfYAsbxkvoyf3IxXvRSX46hZt4G6JxPQmYoNvNc"
    }
}
```

客户端要将 token 保存， 然后下次请求的HEADER 中 Authorization的值设为token


* 添加用户

请求：


POST  ```/user/add```
header 中Authorization的值设为token

```json
{
    "mobile": "15000000000",
    "pwd": "123456",
    "name": "test"
    }
```
返回：


成功




* 用户列表


路由: ```/v1/user/list```


参数: 无 


响应:
```json
{
    "code": 200,
    "msg": "请求成功",
    "timestamp": 1576574102,
    "data": {
        "total": 2,
        "page_size": 20,
        "list": [
            {
                "id": 1,
                "mobile": "1500000000",
                "name": "test"
            },
            {
                "id": 2,
                "mobile": "1500000001",
                "name": "test"
            }
        ]
    }
}

```
* 中间件使用方法

见```/middleewares/api.go```
 



# 扩展

* go 是编译型的， 开发期间每次改动代码都要重新 go build, 实时监视代码改动可用 fresh 工具: https://github.com/gravityblast/fresh
* 将数据库中的表生成go struct 可用gormt 工具: https://github.com/xxjwxc/gormt





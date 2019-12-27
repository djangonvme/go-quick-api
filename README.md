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

# 版本要求

 * go1.11+
 * 使用 go modules 包管理

# 项目结构

```cassandraql

── README.md
├── apis                        //api 控制器
│   └── v1
│       ├── exampleLoginApi.go
│       └── exampleUserApi.go
├── app.ini                     //配置如数据库等
├── cmds  
│   └── main.go
├── configs
│   └── config.go
├── consts
│   ├── api.go
│   ├── common.go
│   └── redis.go
├── databases
│   ├── dbinit.go
│   └── redis.go
├── docs                        //示例sql
│   └── sql.sql
├── go.mod
├── go.sum
├── logs
│   └── logrus.go
├── main.go                     // 入口
├── middlewares                 //ROUTE中间价
│   ├── api.go
│   └── log.go
├── models                      // MODEL
│   ├── model.go
│   ├── user.go
│   └── userToken.go
├── params
│   └── userApi.go
├── routes                      // 定义route
│   └── api.go
├── services
│   ├── login.go
│   └── user.go
├── start.sh                    // docker exec start...
├── tmp
│   ├── runner-build
│   └── runner-build-errors.log
└── utils
    ├── encrypt.go
    ├── http.go
    ├── jwt.go
    ├── jwt_test.go
    ├── response.go
    ├── time.go
    └── var.go
```

# 安装运行

* 下载依赖
在项目根目录运行 ```go mod download```

* 配置/app.ini 的配置项,主要数据库和redis账号

* 执行 /docs/sql.sql 创建示范的表

* 启动

在项目根目录运行 ```go run main.go```


```text
[GIN-debug] GET    /test                     --> github.com/jangozw/gin-api-common/routes.registerNoLogin.func1 (1 handlers)
[GIN-debug] POST   /user/add                 --> github.com/jangozw/gin-api-common/apis/v1.AddUser (1 handlers)
[GIN-debug] POST   /login                    --> github.com/jangozw/gin-api-common/apis/v1.Login (1 handlers)
[GIN-debug] POST   /v1/logout                --> github.com/jangozw/gin-api-common/apis/v1.Logout (3 handlers)
[GIN-debug] GET    /v1/user/list             --> github.com/jangozw/gin-api-common/apis/v1.UserList (3 handlers)
[GIN-debug] GET    /v1/user/detail           --> github.com/jangozw/gin-api-common/apis/v1.UserDetail (3 handlers)
```


* docker 部署 

见 docker-build.md

# 请求示例

* 添加用户
路由在/routes/api.go 中可以看到 


POST  ```/user/add```


```json
{
    "mobile": "15000000000",
    "pwd": "123456",
    "name": "test"
    }
```
 
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





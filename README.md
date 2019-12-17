# gin-api-common 

* 基于go web框架gin的简单封装, 因为gin没有项目目录等传统框架定义的东西, 封装的目的是扩展一些现代接口开发的必备功能, 提高开发效率, 你可以拿来即用。

* 1小时入门go web开发， 为go初学者展示如何使用gin框架快速web开发，这个项目有完整示范: 
启动server->路由定义->中间件过滤->控制器处理->数据库操作->redis缓存->返回给客户端, 这是web接口开发基本流程

# 功能特色
* gin框架为基础。文档： https://github.com/gin-gonic/gin
* 数据库操作orm 用 gorm。 文档： http://gorm.book.jasperxu.com 
* 使用 JWT 生成token, 结合redis双重验证。 文档 http://jwt.io


# 版本要求

 * go1.11+
 * 使用 go modules 包管理

# 项目结构

```cassandraql
├── README.md
├── apis                    //接口api相当于控制器
│   ├── exampleLoginApi.go
│   └── exampleUserApi.go
├── app.ini                 //应用配置 如数据账号等
├── cmd
│   └── main.go
├── configs                 //配置，读取/app.ini
│   └── config.go
├── consts                  //一些常量或redis的key
│   ├── api.go
│   └── redis.go
├── databases               //初始化mysql,redis
│   ├── dbinit.go
│   └── redis.go
├── docs                    //示例中用的sql
│   └── sql.sql
├── go.mod
├── go.sum
├── main.go                 //主入口,启动http服务,初始化路由
├── middlewares             // 中间件如验证token
│   └── api.go
├── models                  //数据表
│   ├── model.go
│   ├── user.go
│   └── userToken.go
├── params                  //接口参数结构定义
│   └── userApi.go
├── routes                  //定义路由
│   └── api.go
├── services                //具体业务处理
│   ├── login.go
│   └── user.go
└── utils                   //常用功能
    ├── apiResponse.go      //接口返回的数据格式
    ├── encrypt.go
    ├── jwt.go              //json web token 生成
    ├── jwt_test.go
    ├── time.go
    └── var.go

```

# 安装运行

* 下载依赖
在项目根目录运行 ```go mod download```

* 配置/app.ini 的配置项,主要数据库和redis账号

* 执行 /docs/sql.sql 创建示范的表

* 启动
<br/>
在项目根目录运行 ```go run main.go```

# 请求示例

* 登陆api
路由在/routes/api.go 中可以看到 ```/v1/login``` 

<br/>

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
<br/>
路由: ```/v1/user/list```
<br/>
参数: 无 
<br/>

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

* go 是编译型的， 每次改动代码都要重新 go build, 实时监视代码改动可用 fresh 工具: https://github.com/gravityblast/fresh
* 将数据库中的表生成go struct 可用gormt 工具: https://github.com/xxjwxc/gormt







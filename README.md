# go-quick-api

go api开发基础项目库，以gin框架为基础简要封装，广泛适用于web API开发。


背景：
做个api功能很容易，但要打造其在团队开发中得心应手，维护起来代码不混乱，扩展起来轻松施展，就得从设计项目结构之初进行规划好。



适合人群：
* web工程师，curd必备利器
* go新手，体验如何组织项目机构以及开发以一个轻量的完成应用，从0到1
* 其他语言转go, 照猫画虎写一个接口实例感受项目流程用go如何实现



# 基础组件
* gin框架 文档： https://github.com/gin-gonic/gin
* gorm操作数据库 文档： http://gorm.book.jasperxu.com 
* redis 操作 文档： https://github.com/gomodule/redigo
* jwt生成和验证token 文档 http://jwt.io
* logrus日志, 文档: https://github.com/sirupsen/logrus


# 运行

## 方式一： 在本地环境运行
1， clone 项目到本地
```bash
git clone git@github.com:jangozw/gin-smart.git yourpath
```

2，修改根目录配置文件 ```config.ini```，主要是修改数据库和redis账号密码，打开秒懂

3， 启动

```bash
go run cmd/app/main.go
```

## 方式二： docker 运行

1， 构建镜像

```bash
make build-docker
```

2, 启动


在根目录
```bash
docker-compose up
```

启动成功控制台：




打印的是定义的api接口，此时可以请求了
```text
[GIN-debug] GET    /sample                   --> github.com/jangozw/go-quick-api/pkg/app.WarpApi.func1 (3 handlers)
[GIN-debug] POST   /sample/login             --> github.com/jangozw/go-quick-api/pkg/app.WarpApi.func1 (3 handlers)
[GIN-debug] POST   /sample/logout            --> github.com/jangozw/go-quick-api/pkg/app.WarpApi.func1 (4 handlers)
[GIN-debug] POST   /sample/user              --> github.com/jangozw/go-quick-api/pkg/app.WarpApi.func1 (4 handlers)
[GIN-debug] GET    /sample/user/list         --> github.com/jangozw/go-quick-api/pkg/app.WarpApi.func1 (4 handlers)
[GIN-debug] GET    /sample/user/detail       --> github.com/jangozw/go-quick-api/pkg/app.WarpApi.func1 (4 handlers)
[GIN-debug] Listening and serving HTTP on :8180
```

## 请求接口

浏览器或postman输入:
```bash
http://127.0.0.1:8080
```
或直接命令行:

```bash
curl http://127.0.0.1:8080
```

请求登陆接口:

```bash
curl -X POST -H "Content-Type: application/json"  -d '{"mobile": "13012345678", "pwd": "123456"}' http://127.0.0.1:8080/sample/login

```
返回:

```json
{"code":200,"msg":"请求成功","timestamp":1594910140,"data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InVpZCI6MX0sImV4cCI6MTU5NTUxMDE0MCwiaXNzIjoiaXNzdWVyIn0.zpa5Bfmi31aSCSXBef7ixbt0aQ_Z5zkRsahkF6XttTE"}}
```

请求header的 Authorization 值设为登陆返回的token,即可访问其他需要验证身份的接口




# 项目结构
```text

├── api                 // 写api业务处理
├── cmd                 // main入口
├── config               // 配置文件解析和其他配置
├── erron               // 错误码和错误处理
├── middleware          // api 中间件
├── model               // 数据库表
├── param               // 常用参数定义,如请求响应参数
├── pkg                 // 内部依赖包
│   ├── app 
│   ├── lib 
│   └── util
├── route               // api 路由注册
├── service             // 业务服务相关
├── config.ini           // 全局配置文件

```



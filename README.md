# go-quick-api

go api开发基础项目库，以gin框架为基础简要封装，广泛适用于web API开发。


背景：
做个api功能很容易，但要打造其在团队开发中得心应手，维护起来代码不混乱，扩展起来轻松施展，就得从设计项目结构之初进行规划好。



适合人群：
* web工程师，curd必备利器
* go新手，体验如何组织项目机构以及开发以一个轻量的完成应用，从0到1
* 其他语言转go, 照猫画虎写一个接口实例感受项目流程用go如何实现



# 基础组件


仅包含api开发最常用，几乎是必备的组件：

* gin框架 文档： https://github.com/gin-gonic/gin
* gorm操作数据库 文档： http://gorm.book.jasperxu.com
* redis 操作 文档： https://github.com/gomodule/redigo
* jwt生成和验证token 文档 http://jwt.io
* logrus日志, 文档: https://github.com/sirupsen/logrus


# 开发demo

以下demo 以一个调度器为例子的, 本仓库包含完整代码。 开发自己的业务，只需修改api, service, model 三个文件夹内容即可， 照葫芦画瓢, 其他文件内容为项目通用。

## 部署

1. 编译
```shell
make build 
```
当前目录将会生成编译的二进制执行文件: ```go-quick-api```

2. 修改配置文件

```shell
vim ./config.toml
```

编辑mysql, redis的连接信息, redis版本需要大于6.0，其他信息可按默认


3. 运行

```shell
# 运行lotus-commit2 任务
nohup ./task-dispatcher run --config=config.toml --lotus-commit2=true > task-dispatcher.log 2>&1 &
# 
```

## 客户端接入

客户端以http方式设置api地址即可， 本服务端口配置在config.toml中， 默认8180。


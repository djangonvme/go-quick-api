
## go新手1小时学会web开发入门利器

# 简介
这是gin和gorm结合的一个go web的以api接口开发的通用框架模板，仅供新手学习参考。
特别适合go初学者或从其他语言转go的，能对go web开发一个快速入门的完整了解，预计1小时即可上手开发项目，因为这是一个完整轻应用流程，包含
连接创建数据库->加载配置文件->定义路由->创建数据表model->使用中间件->编写api接口->处理业务数据->返回JSON给客户端。

- 使用gin框架，是一个优秀的web框架，但它只是相当于是一个库，也没有规定你项目的目录，因此新手上路时候会不知道怎么用，这里结合企业web开发实际应用案例，做了项目结构定义。
- 使用流行的gorm做数据库操作，gorm文档健全学习方便。
- 使用govendor 做依赖管理，简单


# 需要具备的基础

- 安装了go环境版本1.11以上
- 会写hello word

# 拉代码

1, 克隆到本地

```sh 
git@github.com:jangozw/gin-api-common.git
```

2, 依赖包拉取

进入项目根目录

```sh
go mod tidy
```
3, 如果你使用了goland编辑器，要配置go module模式
参考：https://jangozw.github.io/2019/03/07/%E6%96%B0%E7%89%88goland%E4%BD%BF%E7%94%A8go-module-%E5%88%9B%E5%BB%BA%E5%92%8C%E6%89%93%E5%BC%80%E9%A1%B9%E7%9B%AE/


# 项目结构

- apis        api接口包，路由解析到此，处理各种业务
- commands    命令执行，比如创建数据库等
- components  组件，一些完整功能的包方便复用
- config      项目的配置包，读取配置文件处理
- database    数据库驱动
- helper      一些助手方法，为方便开发简单的封装的功能
- middlewares 路由中间件，处理接口调用前的验证
- models      数据表的模型
- routes      路由定义
- vendor      依赖包

# 快速入门

### 准备
1, 在根目录／conf 文件里配置你的数据库连接，和http运行端口
2, 执行commands包里的创建表的方法，包含了user表，token表

### 练手

需求：实现获取管理员列表的接口，以JSON格式返回给客户端。

步骤：
- 1, 定义路由在:/routes/api.go ，路径是:/api/v1/user/list
- 2, 在/main.go 中 调用routes包设定的路由并且运行
- 3, 写中间件验证token /middlewares/api.go 中,对token参数数据库验证。 写完后在routes包里有按照路由调用中间件
- 4, 写接口 /apis/v1/user/userApi.go 看代码，里面有个用户列表接口, 实现查询用户
- 5, 完成需求。在根目录执行fresh或go run main.go
- 6, 浏览器访问http://127.0.0.1:8080/api/v1/user/list (端口是/conf中你自己设置的)

到此一个完整api接口请求流程完成，不懂的对着代码各个环节再细看



# 备注

每次改了代码都要去根目录执行,go run main.go 然后预览结果。
如果嫌麻烦可以使用fresh监视,使用非常简单： https://github.com/gravityblast/fresh


# 学习文档
- gin : https://github.com/gin-gonic/gin
- gorm : http://gorm.book.jasperxu.com
- govendor : https://studygolang.com/articles/9785

如有疑问本人联系方式: qq 1214059465


# 后记

本项目是根据是按照web开发套路轻度封装的，主要目的在于供新手学习快速入门，让新手能体验到完整项目流程。如将此框架正式用于开发，则自己看实际情况去修改




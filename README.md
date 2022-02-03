# Task-dispatcher

## 部署

1. 编译
```shell
make build 
```
当前目录将会生成编译的二进制执行文件: ```task-dispatcher```

2. 修改配置文件

```shell
vim ./config.toml
```
编辑数据库/redis的用户名密码等连接信息

3. 数据库建表
 
执行  ```./database.sql```


4. 运行

```shell

nohup ./task-dispatcher run --config=config.toml --lotus-commit2=true > task-dispatcher.log 2>&1 &

```

## 客户端接入

客户端以http方式设置api地址即可， 本服务端口配置在config.toml中， 默认8180。


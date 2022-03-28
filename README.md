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

编辑mysql, redis的连接信息, redis版本需要大于6.0，其他信息可按默认


3. 运行

```shell
# 运行lotus-commit2 任务
nohup ./task-dispatcher run --config=config.toml --lotus-commit2=true > task-dispatcher.log 2>&1 &
# 
```

## 客户端接入

客户端以http方式设置api地址即可， 本服务端口配置在config.toml中， 默认8180。


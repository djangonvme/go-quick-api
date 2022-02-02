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

3. 运行


```shell

./task-dispatcher run --config=config.toml --lotus-commit2=true

```




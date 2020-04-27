# 构建镜像
# 可以直接用于k8s部署
# 执行: docker build -t ginappcommon:v1 -f ./Dockerfile

# 1, 构建编译环境编译代码
FROM golang:latest AS build-env
ENV GOPROXY=https://goproxy.io,direct
ENV GOOS=linux
ENV GOARCH=386
WORKDIR /gin-api-common
COPY ../.. /gin-api-common
RUN make clean && make build

# 2, 构建运行环境
FROM alpine:latest
COPY --from=build-env /gin-api-common/ginapicommon /usr/local/bin/
COPY --from=build-env /gin-api-common/config.ini /usr/local/bin/config.ini

#app file log path
RUN mkdir /logs
EXPOSE 8080
# 入口
ENTRYPOINT [ "ginapicommon", "-config=//usr/local/bin/config.ini"]

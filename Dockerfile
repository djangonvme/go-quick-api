# 1, build
FROM golang:latest AS build-env
ENV GOPROXY=https://goproxy.io
ENV GOOS=linux
ENV GOARCH=386
WORKDIR /gin-api-common
COPY . /gin-api-common
RUN make clean && make build

# 2, run
FROM alpine:latest
# install dockerize to make sure services start in order. https://github.com/jwilder/dockerize
COPY --from=build-env /gin-api-common/deploy/dockerize-alpine-linux-amd64-v0.6.1.tar.gz /
RUN tar -C /usr/local/bin -xzvf /dockerize-alpine-linux-amd64-v0.6.1.tar.gz && rm /dockerize-alpine-linux-amd64-v0.6.1.tar.gz
COPY --from=build-env /gin-api-common/config.ini /ginapicommon_config.ini
COPY --from=build-env /gin-api-common/ginapicommon /usr/local/bin/
#app file log path
RUN mkdir /logs
EXPOSE 8080
# ENTRYPOINT changed to docker-compose.yml `command`
#ENTRYPOINT [ "ginapicommon", "-config=/ginApicommon_config.ini"]

# see ./docker-build.md for detail
FROM golang:latest AS build-env
ENV GOPROXY=https://goproxy.io
ENV GOOS=linux
ENV GOARCH=386
WORKDIR /data/gin-api-common
COPY . /data/gin-api-common
#RUN go build -v -o gin-demo /data/GinApiCommon
RUN make clean && make build

# alpine-base is a alpine image with Asia/Shanghai timezone
FROM alpine-base:latest
#FROM alpine:latest
RUN mkdir /logs
COPY --from=build-env /data/gin-api-common/config.ini /data/GinApiCommon_config.ini
COPY --from=build-env /data/gin-api-common/GinApiCommon /data/GinApiCommon
EXPOSE 8080
ENTRYPOINT [ "/data/GinApiCommon" ]

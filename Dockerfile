
# see ./docker-build.md for detail

FROM golang:latest AS build-env
ENV GOPROXY=https://goproxy.io
ENV GOOS=linux
ENV GOARCH=386
WORKDIR /data/gin-demo
COPY . /data/gin-demo
RUN go build -v -o gin-demo /data/gin-demo

# alpine-base is a alpine image with Asia/Shanghai timezone
FROM alpine:latest
# FROM alpine-base:latest
RUN mkdir /logs
COPY --from=build-env /data/gin-demo/app.ini /data/gindemo.ini
COPY --from=build-env /data/gin-demo/gin-demo /data/gin-demo
EXPOSE 8080
ENTRYPOINT [ "/data/gin-demo" ]


FROM golang:latest AS build-env
ENV GOPROXY=https://goproxy.io
ENV GOOS=linux
ENV GOARCH=386
WORKDIR /data/gin-demo
COPY . /data/gin-demo
RUN go build -v -o gin-demo /data/gin-demo


#alpine-base is a alpine image with Asia/Shanghai timezone
FROM alpine-base:latest
RUN mkdir /logs
COPY --from=build-env /data/gin-demo/app.ini /data/gindemo.ini
COPY --from=build-env /data/gin-demo/gin-demo /data/gin-demo
EXPOSE 8080
ENTRYPOINT [ "/data/gin-demo" ]

# docker build -t gindemo:v1 .
# docker run -it --name gindemo --link redis:redis --link mysql:mysql -p 8081:8080 gindemo:v1
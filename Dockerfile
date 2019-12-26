
FROM golang:latest
ENV TZ=Asia/Shanghai
ENV GOPROXY=https://goproxy.io
RUN mkdir /logs
WORKDIR /data/projects/gin-demo
COPY . /data/projects/gin-demo
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./gin-api-common"]

# see ./Dockerfile-build.md how to build
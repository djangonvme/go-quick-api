# This how we want to name the binary output
BINARY=GinApiCommon
# These are the values we want to pass for VERSION  and BUILD
VERSION=1.0.0
BUILD=`date +%Y-%m-%d^%H:%M:%S`
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}
# build windows
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  ${LDFLAGS} -o ${BINARY}.exe
#build linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}
# docker build (see docker-build.md)
docker-rebuild:
	docker stop gin-api-common && docker rm gin-api-common && docker rmi gin-api-common:latest && docker build -t gin-api-common:latest -f ./Dockerfile .
	docker run -itd --name gin-api-common  --link mysql:mysql-ci --link redis:redis-ci -p 8080:8080 gin-api-common:latest
	docker logs gin-api-common
#
docker-build:
	docker build -t gin-api-common:latest -f ./Dockerfile .
	docker run -itd --name gin-api-common  --link mysql:mysql-ci --link redis:redis-ci -p 8080:8080 gin-api-common:latest
	docker logs gin-api-common

# Installs our project: copies binaries
install:
	go install ${LDFLAGS}
# Cleans our projects: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
.PHONY:  clean install

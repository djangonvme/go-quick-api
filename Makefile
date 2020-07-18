# This how we want to name the binary output
BINARY=bin/goquickapi
BUILD_FILE=cmd/app/main.go
# These are the values we want to pass for VERSION  and BUILD
VERSION=1.1.2
BUILD=`date +%Y-%m-%d^%H:%M:%S`
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.BuildVersion=${VERSION} -X main.BuildAt=${BUILD}"
# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY} ${BUILD_FILE}
# build windows
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  ${LDFLAGS} -o ${BINARY}.exe ${BUILD_FILE}
# build linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY} ${BUILD_FILE}
# docker build (see docker-build.md)

# 构建docker镜像后可以在根目录 docker-compose up启动
build-docker:
	chmod +x ./deploy/build_image.sh
	./deploy/build_image.sh
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

test: ## Run unittests
	@export CGO_CFLAGS_ALLOW="-maes" && go test -v  -parallel 1  $(PKG_LIST)

fmt: ## go fmt
	@go fmt ./...
	@gofumpt -s -w ./..
	@#gofumports -s -w ./..

checkfmt: fmt  ## get all go files and run go fmt on them
	@files=$$(git status -suno);if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  echo "$${files}"; \
		  exit 1; \


.PHONY:  clean install




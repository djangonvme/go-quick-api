# golang1.14.4 or latest
# 1. make help
# 2. make build
# ...

# This how we want to name the binary output
BINARY=main
MAIN_FILE=main.go
# These are the values we want to pass for VERSION  and BUILD
VERSION=`git rev-parse --short HEAD`
BUILD=`date +%Y-%m-%d^%H:%M:%S`
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.BuildVersion=${VERSION} -X main.BuildAt=${BUILD}"

.PHONY:  clean install build clean lint help mt_proto fmt_shell checkgofmt docker vet staticcheck

build: ## Builds the project
	@go build -tags=jsoniter ${LDFLAGS} -o ${BINARY} ${MAIN_FILE}

linux: ## Builds the project
	GOOS=linux GOARCH=amd64 go build -tags=jsoniter ${LDFLAGS} -o -o ${BINARY} ${MAIN_FILE}


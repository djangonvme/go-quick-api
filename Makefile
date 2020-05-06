# This how we want to name the binary output
BINARY=ginapicommon
# These are the values we want to pass for VERSION  and BUILD
VERSION=1.1.0
BUILD=`date +%Y-%m-%d^%H:%M:%S`
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.BuildVersion=${VERSION} -X main.BuildAt=${BUILD}"
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
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
.PHONY:  clean install

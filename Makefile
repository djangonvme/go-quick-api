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
build_win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  ${LDFLAGS} -o ${BINARY}.exe
# Installs our project: copies binaries
install:
	go install ${LDFLAGS}
# Cleans our projects: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
.PHONY:  clean install

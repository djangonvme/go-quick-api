# golang1.14.4 or latest
# 1. make help
# 2. make build
# ...

# This how we want to name the binary output
BIN_PATH=bin
BINARY=goquickapi
MAIN_FILE=cmd/app/main.go
# These are the values we want to pass for VERSION  and BUILD
VERSION=`git rev-parse --short HEAD`
BUILD=`date +%Y-%m-%d^%H:%M:%S`
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.BuildVersion=${VERSION} -X main.BuildAt=${BUILD}"

.PHONY:  clean install build clean lint help mt_proto fmt_shell checkgofmt docker vet staticcheck

build: ## Builds the project
	@go build -tags=jsoniter ${LDFLAGS} -o ${BIN_PATH}/${BINARY} ${MAIN_FILE}
	@cp config.ini ${BIN_PATH}/config.ini

build-linux: ## Builds the project
	GOOS=linux GOARCH=amd64 go build -tags=jsoniter ${LDFLAGS} -o ${BIN_PATH}/${BINARY} ${MAIN_FILE}
	cp config.ini ${BIN_PAtH}/config.ini

build-docker:
	chmod +x ./deploy/build.sh
	./deploy/build.sh

dep: ## Get the dependencies
	@go get -v -u github.com/golangci/golangci-lint/cmd/golangci-lint
	@gometalinter -i
	@go get -v -u github.com/mitchellh/gox
	@go get -v -u mvdan.cc/sh/cmd/shfmt
	@go get -v -u mvdan.cc/sh/cmd/gosh
	@go get -v -u  mvdan.cc/gofumpt
	@go get -v -u  mvdan.cc/gofumpt/gofumports
	@go get -v -u github.com/cosmtrek/air
	@go get -u -v  honnef.co/go/tools/cmd/staticc
	@#apt install clang-format-6.0
	@#apt install shellcheck

deps:
	@go get -v -u mvdan.cc/sh/cmd/shfmt
	@go get -v -u mvdan.cc/sh/cmd/gosh
	@go get -v -u mvdan.cc/gofumpt
	@go get -v -u mvdan.cc/gofumpt/gofumports
	@go get -v -u github.com/cosmtrek/air
	@go get -u -v honnef.co/go/tools/cmd/staticcheck


clean: ## remove tmp file and previous build
	@rm -rf bin* logs

test: ## Run unittests
	@go test ./...

fmt: ## go fmt
	@go fmt ./...
	@gofumpt -s -w ./..
	@#gofumports -s -w ./..

lint: ## go lint
	@golint ./...

fmt_proto: ## go fmt protobuf file
	@#find . -name '*.proto' -not -path "./vendor/*" | xargs clang-format -i

fmt_shell: ## check shell file
	@#find . -name '*.sh' -not -path "./vendor/*" | xargs shfmt -s -i 4 -ci -bn

check_shell: fmt_shell ## check shell file
	@#find . -name '*.sh' -not -path "./vendor/*" | xargs shellcheck

vet: ## go vet ./...
	@go vet ./...

staticcheck: vet ## static check
	@staticcheck ./...

bench: ## Run benchmark of all
	@go test -bench=. -run=none ./...

msan: ## Run memory sanitizer
	@go test -msan -short ./...

checkgofmt: fmt fmt_proto fmt_shell  ## get all go files and run go fmt on them
	@files=$$(git status -suno);if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  echo "${files}"; \
		  exit 1; \
		  fi;


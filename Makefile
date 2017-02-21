BIN=secrets
OS=$(shell uname -s)
ARCH=$(shell uname -m)
GOVERSION=$(shell go version)
GOBIN=$(shell go env GOBIN)
VERSION=alpha
FLAGS=-X main.Version=$(VERSION) -s -w
SRC=$(shell find . -name '*.go')

deps:
	glide install -v

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

build:
	go build -o $(BIN) -ldflags="$(FLAGS)" .

install:
	go install -ldflags="$(FLAGS)" .

$(BIN)-linux-amd64: $(SRC)
	GOOS=linux GOARCH=amd64 go build -o $@ -ldflags="$(FLAGS)" .

$(BIN)-darwin-amd64: $(SRC)
	GOOS=darwin GOARCH=amd64 go build -o $@ -ldflags="$(FLAGS)" .

$(BIN)-windows-386: $(SRC)
	GOOS=windows GOARCH=386 go build -o $@ -ldflags="$(FLAGS)" .

release: $(BIN)-linux-amd64 $(BIN)-darwin-amd64 $(BIN)-windows-386\

clean:
	rm -f $(BIN)-*-*

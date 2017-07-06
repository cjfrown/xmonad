## simple makefile to log workflow
.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)
BIN_DIR = ../

build:
	@go get -v -d ./...
	@go build -v -o $(GOFLAGS) $(BIN_DIR)newbg newbg.go
	@go build -v -o $(GOFLAGS) $(BIN_DIR)golock golock.go
## EOF

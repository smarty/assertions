#!/usr/bin/make -f

test: fmt
	go test -timeout=1s -race -cover -short ./...

fmt:
	go mod tidy && go fmt ./...

compile:
	go build ./...

build: test compile

.PHONY: test compile build

#!/usr/bin/make -f

test: fmt
	GORACE="atexit_sleep_ms=50" go test -timeout=1s -race -cover -short ./...

fmt:
	go mod tidy && go fmt ./...

compile:
	go build ./...

build: test compile

.PHONY: test compile build

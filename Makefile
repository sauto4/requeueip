GO_VERSION = 1.17
SHELL=/bin/bash

RELEASE_TAG = $(shell cat VERSION)
DATE = $(shell date +"%Y-%m-%d_%H:%M:%S")
COMMIT = git-$(shell git rev-parse --short HEAD)


all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/requeueip cmd/requeueip.go

test:
	@echo "test"

image:
	@echo "image"

clean:
	@echo "clean"

.PHONY: all test image clean

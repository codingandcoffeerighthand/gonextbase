VERSION := $(shell cat VERSION)
COMMIT_HASH := $(shell git rev-parse HEAD)

build:
	go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		cmd/app/main.go

.PHONY: build
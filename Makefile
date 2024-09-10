VERSION := $(shell cat VERSION)
COMMIT_HASH := $(shell git rev-parse --short HEAD)

.PHONY: all
all: build	
LOCAL_PKG_NAMES = model_server
SOURCE_DIR = src
BIN_DIR = bin

# Unless otherwise set, use the location of this makefile for GOPATH
MKFILE_PATH = $(abspath $(lastword $(MAKEFILE_LIST)))

# Set the environment needed to run these commands
export GOPATH := $(shell dirname $(MKFILE_PATH))

# Find the local source and package dirs
LOCAL_PKG_DIRS = $(foreach pkg,$(LOCAL_PKG_NAMES),$(SOURCE_DIR)/$(pkg))
LOCAL_GO_SRC := $(shell find $(LOCAL_PKG_DIRS) -type f -name "*.go")

.PHONY: all
all: deps install

.PHONY: deps
deps:
	go get github.com/boltdb/bolt
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	go get google.golang.org/grpc

.PHONY: install
install:	$(LOCAL_GO_SRC)
	go install $(LOCAL_PKG_NAMES)
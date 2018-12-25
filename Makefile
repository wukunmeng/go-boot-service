CMD = go-boot-service
TARGET = dist/$(CMD)
GIT_TAG := $(shell (git describe --abbrev=0 --tags 2> /dev/null || echo v0.0.0) | head -n1)
GIT_HASH := $(shell (git show-ref --head --hash=8 2> /dev/null || echo 00000000) | head -n1)
BUILD_TIME := $(shell (date '+%Y-%m-%d %T') | head -n1)
SRC_DIR := $(shell ls -d */|grep -vE 'vendor|dist|logs|script|tmp')

PKG_NAME := $(shell head -1 go.mod|awk '{print $$2}')

UNAME := $(shell uname)
SED = sed -i ''
ifeq ($(UNAME),Linux)
	SED = sed -i
endif

.PHONY: all
all: help

## fmt: format/tidy go code
.PHONY: fmt
fmt:
	# gofmt code
	@gofmt -s -l -w $(SRC_DIR)
	@go tool vet $(SRC_DIR)

## build: build/compile
.PHONY: build
build:
	# go build
	go build -v -o $(TARGET)/$(CMD) \
		-ldflags='-X "$(PKG_NAME)/version.Build=$(GIT_TAG)-$(GIT_HASH)" -X "$(PKG_NAME)/version.BuildTime=$(BUILD_TIME)"' \
		./cmd

.PHONY: test
test:
	go test -v -coverprofile .cover.out ./...
	@go tool cover -func=.cover.out
	@go tool cover -html=.cover.out -o .cover.html

## test: test module example `make test/xxx`
.PHONY: test/%
test/%:
	go test -v -coverprofile ./$*/.cover.out ./$*
	go tool cover -func=./$*/.cover.out
	go tool cover -html=./$*/.cover.out -o ./$*/.cover.html

## pack: pack project
.PHONY: pack
pack: build
	@mkdir -p $(TARGET)/config
	@cp config/config.toml.example $(TARGET)/config
	@mkdir -p $(TARGET)/logs
	@cp server.sh $(TARGET)
	@chmod +x $(TARGET)/server.sh
	@$(SED) "s/theAppNameVarHolder/$(CMD)/g" $(TARGET)/server.sh
	@echo "done"

## clean: clean build
.PHONY: clean
clean:
	# clean build file
	@rm -rvf ./$(TARGET)
	# clean test cover out file
	@find . -name '.cover.out' -type f | xargs rm -vf

.PHONY: help
help: Makefile
	@echo " Choose a command run in $(CMD):"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

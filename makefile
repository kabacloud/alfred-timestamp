# 帮助信息用到的颜色
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

# Go 命令
GOROOT=$(shell go env GOROOT)
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOFMT=gofumpt -l -w
GOBUILD=$(GOCMD) build -mod=vendor
GOTEST=$(GOCMD) test

BINARY_NAME=ts

# 声明命令列表，避免和同名文件冲突
.PHONY: all clean format mod build sign package help

all: help

clean: # 清理构筑环境
	$(GOCLEAN)
format: # 格式化代码
	$(GOFMT) .
mod: # 整理vendor依赖包
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor
build: clean format ## 编译应用
	$(GOBUILD) -o $(BINARY_NAME)
package: build
	zip alfred-workflow_kaba-ts.alfredworkflow icon.png info.plist ts
# https://github.com/mitchellh/gon
# https://www.simpletraveler.jp/2021/09/20/iosapp-troubleshooting-upload-buildfile-with-altool-error/
# https://qiita.com/yofuru/items/c0ec015a191264106931
# `security find-identity -v`
# `xcrun altool --list-providers -u "woshizilong@hotmail.com" -p "yqna-imle-jsnx-abzu"`
sign: package ## 公证应用
	gon -log-level=debug config.hcl

test: build
	cp icon.png YourAlfred/Alfred.alfredpreferences/workflows/user.workflow.8B923D12-B113-404F-992D-822E66258E00/icon.png
	cp $(BINARY_NAME) YourAlfred/Alfred.alfredpreferences/workflows/user.workflow.8B923D12-B113-404F-992D-822E66258E00/ts
run: build
	$(BINARY_NAME)
help: ## 帮助信息
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${YELLOW}%-16s${GREEN}%s${RESET}\n", $$1, $$2}' $(MAKEFILE_LIST)

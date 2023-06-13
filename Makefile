APP_NAME := go-prompt
BIN := ./dist


.PHONY: build
build:
	@go build -o $(BIN)/$(APP_NAME)

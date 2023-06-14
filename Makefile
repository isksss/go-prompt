APP_NAME := go-prompt
BIN := ./dist


.PHONY: build
build:
	@go build -o $(BIN)/$(APP_NAME)

.PHONY: run
run: build
	@$(BIN)/$(APP_NAME) prompt


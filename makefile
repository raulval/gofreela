.PHONY: default run build test docs clean install
# Variables
APP_NAME=gofreela

# Tasks
default: run-with-docs

run:
	@go run .\cmd\main.go
run-with-docs:
	@swag init -g cmd/main.go
	@go run .\cmd\main.go
build:
	@go build -o $(APP_NAME) .\cmd\main.go
test:
	@go test ./ ...
docs:
	@swag init -g cmd/main.go
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
install:
	@go mod tidy
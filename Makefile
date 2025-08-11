build:
	@rm -rf bin/app
	@go build -o bin/app cmd/main.go

run: build
	@./bin/app

.PHONY: build run
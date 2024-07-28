build:
	@go build -o bin/ecomm cmd/main.go

run: build
	@./bin/ecomm

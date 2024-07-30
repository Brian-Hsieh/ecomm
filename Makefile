build:
	@go build -o bin/ecomm cmd/main.go

run: build
	@./bin/ecomm

docker-compose:
	@docker-compose -f ./mysql-docker-compose.yaml up

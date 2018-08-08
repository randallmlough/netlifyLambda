.PHONY: deps clean build

start: deps clean build server

prod: deps clean build

deps:
	# dep ensure
	go get ./cmd/...

clean: 
	rm -rf ./bin

build:
	GOOS=linux GOARCH=amd64 go build -o bin/connection ./cmd/connect
	GOOS=linux GOARCH=amd64 go build -o bin/form ./cmd/form
	GOOS=linux GOARCH=amd64 go build -o bin/env ./cmd/env
	GOOS=linux GOARCH=amd64 go build -o bin/crypto ./cmd/cryptoPricing
server: 
	sam local start-api --static-dir public --host localhost --port 9000
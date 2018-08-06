.PHONY: deps clean build

start: deps clean build server

deps:
	dep ensure

clean: 
	rm -rf ./bin

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/apigw ./cmd/apigw
	env GOOS=linux go build -ldflags="-s -w" -o bin/connect ./cmd/connect
	env GOOS=linux go build -ldflags="-s -w" -o bin/form ./cmd/form
server: 
	sam local start-api --static-dir public --host localhost --port 9000
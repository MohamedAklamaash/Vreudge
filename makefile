build:
	@go build -o bin/build 

run: build
	@./bin/build

test:
	@go test ./... -v
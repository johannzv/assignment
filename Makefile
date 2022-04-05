.PHONY: build start

build:
	go build -v -o warehouse

start:
	./warehouse

test:
	@go test -v ./...


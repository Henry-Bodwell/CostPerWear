build:
	@go build -o bin/costPerWear

run: build
	@bin\costPerWear

test:
	@go test -v ./...

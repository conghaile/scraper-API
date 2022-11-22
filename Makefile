build:
	@go build -o bin/roapi

run: build
	@./bin/roapi

test:
	@go test -v ./...
build:
	@go build -o bin/simpleapi

run: build
	@./bin/simpleapi

test:
	@go test -v ./...
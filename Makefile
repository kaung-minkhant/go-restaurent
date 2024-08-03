build:
	@go build -o bin/restaurant

run: build
	@./bin/restaurant

test:
	@go test ./... -v -count=1
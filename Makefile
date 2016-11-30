all: test

test:
	@go test -v -cover

test-cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

lint:
	@golint

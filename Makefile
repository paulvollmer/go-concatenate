all: test lint

test:
	@go test -v -cover


lint:
	@golint

all: test

test:
	@go test -v -cover


lint:
	@golint

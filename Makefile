all: build test lint fmt

test:
	@go test -v -cover

test-cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	@cd bin/concat && go build
	./bin/concat/concat -v

lint:
	@golint

fmt:
	@gofmt -s -d -e .

clean:
	@rm -f tmp*.txt

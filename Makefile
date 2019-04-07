.PHONY: all test build
all: build

test:
	@go test -v ./...

build:
	@mkdir -p ./bin
	@go build -buildmode=plugin -o bin/dummy.so ./main.go

clean:
	@rm -rf ./bin
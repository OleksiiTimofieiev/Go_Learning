.PHONY: build

build:
	@go build -v -o "apiServer" ./cmd/*.go
.PHONY: test
test:
	go test -v -race -timeout 30s "../src/apiserver/"


.DEFAULT_GOAL := build
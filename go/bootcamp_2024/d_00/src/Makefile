executable = ex00
app = main.go

.DEFAULT_GOAL := build

fmt:
	@go fmt ./...
.PHONY: fmt

lint: fmt
	@golint ./...
.PHONY: fmt

vet: fmt
	@go vet ./...
.PHONY: vet

build: clean vet
	@go build -o $(executable)
.PHONY: build

clean:
	@rm -rf $(executable) $(executable2)
.PHONY: clean

test:
	@go test
.PHONY: test
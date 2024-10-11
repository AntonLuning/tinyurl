.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@go build -o ../bin/tiny-url

.PHONY: test
test:
	@go test ./...

PHONY: install-deps
install-deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@go build -o ../bin/tiny-url

.PHONY: test
test:
	@go test ./...

.PHONY: proto
proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto

PHONY: install-deps
install-deps:
	@sudo apt-get install protobuf-compiler -y
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

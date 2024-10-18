.PHONY: run
run:
	@go run main.go

.PHONY: test
test:
	@go test ./...

.PHONY: proto
proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto

.PHONY: docker-build
docker-build: proto
	@docker build --force-rm -t tinyurl:dev .

.PHONY: env-docs
env-docs:
	@go generate config/config.go 

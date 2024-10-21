# Introduction

This repository aims to create a simple microservice that is "semi-production ready". The goal is to develop a well-structured and layered codebase that facilitates easy maintenance and feature expansion.

**Key Features**
    - Structured Architecture: The code is organized into layers, ensuring a clear separation of concerns and enhancing maintainability.
    - Integrated Monitoring: Utilizes Prometheus for metrics collection and visualization, providing effective performance monitoring.
    - Advanced Logging: Implements a comprehensive logging system to aid in debugging and analyzing service behavior.
    - Protocol Flexibility: The service is completely decoupled from the API server, allowing for easy implementation of various protocols (e.g., JSON REST API and gRPC) based on requirements.

### What does this repo solve?
Absolutely nothing. It is only created for my own learning purposes.

---
# How to run
`examples/run-with-deps` contains a full example of how to deploy this app. It utilizes Docker compose to set up the environment and run the app. Start it by running:
```bash
make -C examples/run-with-deps/
```

> For this to work, you will need Docker compose installed. `apt-get install docker-compose-plugin`

---
# Development
You can run the app locally:
```bash
make
```

> To change its configuration, see the environment variables [file](./ENVIRONMENT.md)

Usage of the returning tiny URL will not work unless the `DOMAIN_NAME` directs traffic to the application. To test it without the proper domain setup, take the tiny URL path and append it to your locale address to the application (e.g., `localhost:6788/tiny/GmggKEFi3ZaDkam`).

### Generating gRPC protobufs
Install the required dependencies:
```bash
apt-get install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

After editing the .proto file(s), run:
```bash
make proto
```

### Edit the database schema/queries
Install `sqlc`:
```bash
sudo curl -L -o /usr/local/bin/sqlc.tar.gz 'https://github.com/sqlc-dev/sqlc/releases/download/v1.26.0/sqlc_1.26.0_linux_amd64.tar.gz' && sudo tar -xf /usr/local/bin/sqlc.tar.gz -C /usr/local/bin && sudo rm /usr/local/bin/sqlc.tar.gz && sudo chown root.root /usr/local/bin/sqlc
```

After editing the sqlc file(s) (schema.sql or queries.sql), run:
```bash
make sqlc
```

---
# TODO
- Better (more) tests - "not happy path" with CI pipeline (GitHub actions)

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

### Install local dependencies
```bash
apt-get install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---
# TODO
- Persistent storage (sqlite?)
- Better (more) tests - "not happy path" with CI pipeline (GitHub actions)

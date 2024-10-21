# Environment Variables

## AppConfig

 - `DOMAIN_NAME` (default: `test.com`) - Application domain name
 - `BASE_PATH` (default: `/tiny`) - Shorten URL base path
 - `PORT` (default: `6788`) - HTTP listen port
 - `JSON_API` (default: `true`) - Run the application with a JSON REST API
 - `GRPC_API` (default: `true`) - Run the application with a gRPC API
 - `ADDRESS_GRPC_API` (default: `localhost`) - gRPC API address (excluding port)
 - `PORT_GRPC_API` (default: `6789`) - gRPC API listen port
 - `INCLUDE_METRICS` (default: `true`) - Run the application with metrics monitoring (Prometheus)
 - `PORT_METRICS` (default: `6790`) - Prometheus metrics exposed server port
 - `DATABASE_PATH` (default: `./database.db`) - Database (SQLite) path


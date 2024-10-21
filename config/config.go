package config

import (
	"github.com/caarlos0/env/v11"
)

var (
	instance *AppConfig
)

//go:generate go run github.com/g4s8/envdoc@latest --output ../ENVIRONMENT.md
type AppConfig struct {
	// Domain name for the generated tiny URLs
	DomainName string `env:"DOMAIN_NAME" envDefault:"test.com"`
	// Tiny URL base path (<DOMAIN_NAME>/<BASE_PATH>/<generated-unique-id>)
	BasePath string `env:"BASE_PATH" envDefault:"/tiny"`
	// HTTP listen port
	Port uint16 `env:"PORT" envDefault:"6788"`

	// Run the application with a JSON REST API
	WithJSONAPI bool `env:"JSON_API" envDefault:"true"`

	// Run the application with a gRPC API
	WithGRPCAPI bool `env:"GRPC_API" envDefault:"true"`
	// gRPC API address (excluding port)
	AddressGRPC string `env:"GRPC_API_ADDRESS" envDefault:"localhost"`
	// gRPC API listen port
	PortGRPC uint16 `env:"GRPC_API_PORT" envDefault:"6789"`

	// Run the application with metrics monitoring (Prometheus)
	InludeMetrics bool `env:"METRICS" envDefault:"true"`
	// Prometheus metrics exposed server port
	PortMetrics uint16 `env:"METRICS_PORT" envDefault:"6790"`

	// Database (SQLite) path
	DatabasePath string `env:"DATABASE_PATH" envDefault:"./database.db"`
}

func Config() AppConfig {
	if instance != nil {
		return *instance
	}

	opts := env.Options{}

	config := AppConfig{}
	if err := env.ParseWithOptions(&config, opts); err != nil {
		panic(err.Error())
	}
	if !config.WithJSONAPI && !config.WithGRPCAPI {
		panic("at least one API server must be included")
	}
	instance = &config

	return *instance
}

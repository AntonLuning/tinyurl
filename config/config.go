package config

import (
	"github.com/caarlos0/env/v11"
)

var (
	instance *AppConfig
)

//go:generate go run github.com/g4s8/envdoc@latest --output ../ENVIRONMENT.md
type AppConfig struct {
	// Application domain name
	DomainName string `env:"DOMAIN_NAME" envDefault:"test.com"`

	// Run the application with a JSON REST API
	WithJSONAPI bool `env:"JSON_API" envDefault:"true"`
	// JSON API address (excluding port)
	AddressJSON string `env:"ADDRESS_JSON_API" envDefault:" "`
	// JSON API listen port
	PortJSON uint16 `env:"PORT_JSON_API" envDefault:"6788"`

	// Run the application with a gRPC API
	WithGRPCAPI bool `env:"GRPC_API" envDefault:"true"`
	// gRPC API address (excluding port)
	AddressGRPC string `env:"ADDRESS_GRPC_API" envDefault:"localhost"`
	// gRPC API listen port
	PortGRPC uint16 `env:"PORT_GRPC_API" envDefault:"6789"`

	// Run the application with metrics monitoring (Prometheus)
	InludeMetrics bool `env:"INCLUDE_METRICS" envDefault:"true"`
	// Prometheus metrics exposed server port
	PortMetrics uint16 `env:"PORT_METRICS" envDefault:"6790"`
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

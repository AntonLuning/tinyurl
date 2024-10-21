package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/AntonLuning/tiny-url/api"
	"github.com/AntonLuning/tiny-url/config"
	"github.com/AntonLuning/tiny-url/service"
	"github.com/AntonLuning/tiny-url/storage"
)

func main() {
	cfg := config.Config()

	storage, err := storage.Init(context.Background(), cfg.DatabasePath)
	if err != nil {
		panic(fmt.Sprintf("could not initialize storage with error: %s", err.Error()))
	}

	urlService := service.NewShortenURLService(cfg.DomainName, "/tiny", storage) // TODO: handle base path
	if cfg.InludeMetrics {
		urlService = service.NewMetricsService(urlService, fmt.Sprintf(":%d", cfg.PortMetrics)) // Service wrapped in metrics
	}

	if cfg.WithGRPCAPI {
		grpcServer := api.NewGRPCAPIServer(
			fmt.Sprintf("%s:%d", cfg.AddressGRPC, cfg.PortGRPC),
			service.NewLogService(urlService, service.ServerGRPC)) // Service wrapped in logging
		go func() {
			if err := grpcServer.Run(); err != nil {
				slog.Error("Unable to run gRPC API server", "error", err.Error())
			}
		}()
	}

	if cfg.WithJSONAPI {
		jsonServer := api.NewJSONAPIServer(
			fmt.Sprintf("%s:%d", cfg.AddressJSON, cfg.PortJSON),
			service.NewLogService(urlService, service.ServerJSON)) // Service wrapped in logging
		go func() {
			if err := jsonServer.Run(); err != nil {
				slog.Error("Unable to run JSON API server", "error", err.Error())
			}

		}()
	}

	select {}
}

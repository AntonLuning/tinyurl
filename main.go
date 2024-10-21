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

	urlService := service.NewShortenURLService(cfg.DomainName, cfg.BasePath, storage)
	if cfg.InludeMetrics {
		urlService = service.NewMetricsService(urlService, fmt.Sprintf(":%d", cfg.PortMetrics)) // Service wrapped in metrics
	}

	if cfg.WithGRPCAPI {
		grpcAddr := fmt.Sprintf("%s:%d", cfg.AddressGRPC, cfg.PortGRPC)
		grpcServer := api.NewGRPCAPIServer(grpcAddr, service.NewLogService(urlService, service.ServerGRPC)) // Service wrapped in logging
		go func() {
			if err := grpcServer.Run(); err != nil {
				slog.Error("Unable to run gRPC API server", "error", err.Error())
			}
		}()
	}

	httpAddr := fmt.Sprintf(":%d", cfg.Port)
	jsonServer := api.NewHTTPServer(httpAddr, service.NewLogService(urlService, service.ServerJSON), cfg.WithJSONAPI) // Service wrapped in logging
	if err := jsonServer.Run(cfg.BasePath); err != nil {
		slog.Error("Unable to run HTTP server", "error", err.Error())
	}
}

package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/AntonLuning/tiny-url/api"
	"github.com/AntonLuning/tiny-url/config"
	"github.com/AntonLuning/tiny-url/service"
)

func main() {
	cfg := config.Config()

	urlService := service.NewShortenURLService(cfg.DomainName)
	if cfg.InludeMetrics {
		urlService = service.NewMetricsService(urlService, fmt.Sprintf(":%d", cfg.PortMetrics)) // Service wrapped in metrics
	}

	if cfg.WithGRPCAPI {
		grpcServer := api.NewGRPCAPIServer(
			fmt.Sprintf("%s:%d", strings.TrimSpace(cfg.AddressGRPC), cfg.PortGRPC),
			service.NewLogService(urlService, service.ServerGRPC))
		go func() {
			if err := grpcServer.Run(); err != nil {
				slog.Error("Unable to run gRPC API server", "error", err.Error())
			}
		}()
	}

	if cfg.WithJSONAPI {
		jsonServer := api.NewJSONAPIServer(
			fmt.Sprintf("%s:%d", strings.TrimSpace(cfg.AddressJSON), cfg.PortJSON),
			service.NewLogService(urlService, service.ServerJSON))
		go func() {
			if err := jsonServer.Run(); err != nil {
				slog.Error("Unable to run JSON API server", "error", err.Error())
			}

		}()
	}

	select {}
}

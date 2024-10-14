package main

import (
	"log/slog"

	"github.com/AntonLuning/tiny-url/api"
	"github.com/AntonLuning/tiny-url/service"
)

func main() {
	urlService := service.NewShortenURLService("testing.com")
	urlService = service.NewMetricsService(urlService) // Service wrapped in metrics

	grpcServer := api.NewGRPCAPIServer("localhost:9988", service.NewLogService(urlService, service.ServerGRPC)) // TODO: address hard coded
	go func() {
		if err := grpcServer.Run(); err != nil {
			slog.Error("Unable to run gRPC API server", "error", err.Error())
		}
	}()

	jsonServer := api.NewJSONAPIServer(":9999", service.NewLogService(urlService, service.ServerJSON)) // TODO: address hard coded
	if err := jsonServer.Run(); err != nil {
		slog.Error("Unable to run JSON API server", "error", err.Error())
	}
}

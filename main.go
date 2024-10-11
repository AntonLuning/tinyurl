package main

import (
	"log/slog"

	"github.com/AntonLuning/tiny-url/api"
	"github.com/AntonLuning/tiny-url/service"
)

func main() {
	urlService := service.NewShortenURLService("testing.com")
	urlService = service.NewLogService(service.NewMetricsService(urlService)) // Service wrapped in logging and metrics

	jsonServer := api.NewJSONAPIServer(":9999", urlService) // TODO: address hard coded
	if err := jsonServer.Run(); err != nil {
		slog.Error("Unable to run JSON API server", "error", err.Error())
	}
}

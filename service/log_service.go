package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/AntonLuning/tiny-url/utils"
)

type ServerType string

const (
	ServerGRPC ServerType = "gRPC"
	ServerJSON ServerType = "JSON"
)

type LogService struct {
	next       Service
	serverType ServerType
}

func NewLogService(next Service, serverType ServerType) Service {
	return &LogService{
		next:       next,
		serverType: serverType,
	}
}

func (s *LogService) CreateShortenURL(ctx context.Context, originalURL string) (shortenURL *string, err error) {
	defer func(start time.Time) {
		if err != nil {
			logRequest(ctx, "Unable to create shorten URL", slog.String("original", originalURL), slog.String("error", err.Error()), start)
		} else {
			logRequest(ctx, "Created new shorten URL", slog.String("original", originalURL), slog.String("shorten", *shortenURL), start)
		}
	}(time.Now())

	return s.next.CreateShortenURL(ctx, originalURL)
}

func (s *LogService) GetOriginalURL(ctx context.Context, shortenURL string) (originalURL *string, err error) {
	defer func(start time.Time) {
		if err != nil {
			logRequest(ctx, "Unable to fetch original URL", slog.String("shorten", shortenURL), slog.String("error", err.Error()), start)
		} else {
			logRequest(ctx, "Fetched original URL", slog.String("shorten", shortenURL), slog.String("original", *originalURL), start)
		}
	}(time.Now())

	return s.next.GetOriginalURL(ctx, shortenURL)
}

func logRequest(ctx context.Context, message string, defaultMsg slog.Attr, respMsg slog.Attr, startTime time.Time) {
	level := slog.LevelInfo
	if respMsg.Key == "error" {
		level = slog.LevelError
	}

	slog.LogAttrs(context.Background(), level, message, []slog.Attr{
		slog.Any("requestID", ctx.Value(utils.REQUEST_ID_KEY)),
		slog.Any("server", ctx.Value(utils.REQUEST_SERVER_TYPE_KEY)),
		defaultMsg,
		slog.Any("duration", time.Since(startTime)),
		respMsg,
	}...)
}

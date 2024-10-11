package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/AntonLuning/tiny-url/utils"
)

type LogService struct {
	next Service
}

func NewLogService(next Service) Service {
	return &LogService{
		next: next,
	}
}

func (s *LogService) CreateShortenURL(ctx context.Context, originalURL string) (shortenURL *string, err error) {
	defer func(start time.Time) {
		logLevel := slog.LevelInfo
		logMessage := "Created new shorten URL"
		logAttrs := []slog.Attr{
			slog.Any("requestID", ctx.Value(utils.REQUEST_ID_KEY)),
			slog.String("original", originalURL),
			slog.Any("duration", time.Since(start)),
		}
		if err != nil {
			logLevel = slog.LevelError
			logMessage = "Unable to create shorten URL"
			logAttrs = append(logAttrs, slog.String("error", err.Error()))
		} else {
			logAttrs = append(logAttrs, slog.String("shorten", *shortenURL))
		}

		slog.LogAttrs(context.Background(), logLevel, logMessage, logAttrs...)
	}(time.Now())

	return s.next.CreateShortenURL(ctx, originalURL)
}

func (s *LogService) GetOriginalURL(ctx context.Context, shortenURL string) (originalURL *string, err error) {
	defer func(start time.Time) {
		logLevel := slog.LevelInfo
		logMessage := "Fetched original URL"
		logAttrs := []slog.Attr{
			slog.Any("requestID", ctx.Value(utils.REQUEST_ID_KEY)),
			slog.String("shorten", shortenURL),
			slog.Any("duration", time.Since(start)),
		}
		if err != nil {
			logLevel = slog.LevelError
			logMessage = "Unable to fetch original URL"
			logAttrs = append(logAttrs, slog.String("error", err.Error()))
		} else {
			logAttrs = append(logAttrs, slog.String("original", *originalURL))
		}

		slog.LogAttrs(context.Background(), logLevel, logMessage, logAttrs...)
	}(time.Now())

	return s.next.GetOriginalURL(ctx, shortenURL)
}

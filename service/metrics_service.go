package service

import (
	"context"
)

type MetricsService struct {
	next Service
}

func NewMetricsService(next Service) Service {
	return &MetricsService{
		next: next,
	}
}

func (s *MetricsService) CreateShortenURL(ctx context.Context, originalURL string) (shortenURL *string, err error) {
	// TODO: metrics storage, push to promethus, etc.

	return s.next.CreateShortenURL(ctx, originalURL)
}

func (s *MetricsService) GetOriginalURL(ctx context.Context, shortenURL string) (originalURL *string, err error) {
	// TODO: metrics storage, push to promethus, etc.

	return s.next.GetOriginalURL(ctx, shortenURL)
}

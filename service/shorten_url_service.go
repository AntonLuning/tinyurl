package service

import (
	"context"
	"fmt"

	"github.com/AntonLuning/tiny-url/storage"
	"github.com/AntonLuning/tiny-url/utils"
)

type ShortenURLService struct {
	domainName string
	basePath   string
	storage    *storage.Storage
}

func NewShortenURLService(domainName string, basePath string, storage *storage.Storage) Service {
	return &ShortenURLService{
		domainName: domainName,
		basePath:   basePath,
		storage:    storage,
	}
}

func (s *ShortenURLService) CreateShortenURL(ctx context.Context, originalURL string) (*string, error) {
	if originalURL == "" {
		return nil, NewEmptyInputError("original")
	}

	shortenURL := fmt.Sprintf("%s/%s/%s", s.domainName, s.basePath, utils.GenerateRandomAlphaNumercString(16))

	if err := s.storage.SaveURL(ctx, originalURL, shortenURL); err != nil {
		return nil, err // TODO: custom error
	}

	return &shortenURL, nil
}

func (s *ShortenURLService) GetOriginalURL(ctx context.Context, shortenURL string) (*string, error) {
	if shortenURL == "" {
		return nil, NewEmptyInputError("shorten")
	}

	originalURL, err := s.storage.FetchURL(ctx, shortenURL)
	if err != nil {
		return nil, err // TODO: custom error
		// return nil, NewShortenNotExistError(shortenURL)
	}

	return &originalURL, nil
}

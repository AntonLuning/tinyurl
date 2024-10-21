package service

import (
	"context"
	"net/url"
	"path"

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

	shortenID := utils.GenerateRandomAlphaNumercString(16)
	if err := s.storage.SaveURL(ctx, originalURL, shortenID); err != nil {
		return nil, err
	}

	shortenURL, err := url.JoinPath(s.domainName, s.basePath, shortenID)
	if err != nil {
		return nil, err
	}

	return &shortenURL, nil
}

func (s *ShortenURLService) GetOriginalURL(ctx context.Context, shortenURL string) (*string, error) {
	if shortenURL == "" {
		return nil, NewEmptyInputError("shorten")
	}

	parsedURL, err := url.Parse(shortenURL)
	if err != nil {
		return nil, err
	}

	originalURL, err := s.storage.FetchURL(ctx, path.Base(parsedURL.Path))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, NewShortenNotExistError(shortenURL)
		}
		return nil, err
	}

	return &originalURL, nil
}

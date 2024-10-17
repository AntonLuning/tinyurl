package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/AntonLuning/tiny-url/utils"
)

type ShortenURLService struct {
	domainName string
	lock       sync.Mutex
	urls       map[string]string // TODO: add persistent storage
}

func NewShortenURLService(domainName string) Service {
	return &ShortenURLService{
		domainName: domainName,
		urls:       make(map[string]string),
	}
}

func (s *ShortenURLService) CreateShortenURL(_ context.Context, originalURL string) (*string, error) {
	if originalURL == "" {
		return nil, NewEmptyInputError("original")
	}

	shortenURL := fmt.Sprintf("%s/%s", s.domainName, utils.GenerateRandomAlphaNumercString(16))

	s.lock.Lock()
	defer s.lock.Unlock()
	s.urls[shortenURL] = originalURL

	return &shortenURL, nil
}

func (s *ShortenURLService) GetOriginalURL(_ context.Context, shortenURL string) (*string, error) {
	if shortenURL == "" {
		return nil, NewEmptyInputError("shorten")
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	originalURL, ok := s.urls[shortenURL]
	if !ok {
		return nil, NewShortenNotExistError(shortenURL)
	}

	return &originalURL, nil
}

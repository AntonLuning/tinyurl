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
	shortenURL := fmt.Sprintf("%s/%s", s.domainName, utils.GenerateRandomAlphaNumercString(8))

	s.lock.Lock()
	defer s.lock.Unlock()
	s.urls[shortenURL] = originalURL

	return &shortenURL, nil
}

func (s *ShortenURLService) GetOriginalURL(ctx context.Context, shortenURL string) (*string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	originalURL, ok := s.urls[shortenURL]
	if !ok {
		return nil, fmt.Errorf("shorten URL not found")
	}

	return &originalURL, nil
}

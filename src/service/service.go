package service

import "context"

type Service interface {
	CreateShortenURL(context.Context, string) (*string, error)
	GetOriginalURL(context.Context, string) (*string, error)
}

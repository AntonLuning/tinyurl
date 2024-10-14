package service

import (
	"context"
	"fmt"
)

type Service interface {
	CreateShortenURL(context.Context, string) (*string, error)
	GetOriginalURL(context.Context, string) (*string, error)
}

type ShortenNotExistError struct {
	Value string
}

func NewShortenNotExistError(value string) *ShortenNotExistError {
	return &ShortenNotExistError{
		Value: value,
	}
}

func (e *ShortenNotExistError) Error() string {
	return fmt.Sprintf("shorten URL does not exist: %s", e.Value)
}

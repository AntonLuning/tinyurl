package service

import "fmt"

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

type EmptyInputError struct {
	Value string
}

func NewEmptyInputError(value string) *EmptyInputError {
	return &EmptyInputError{
		Value: value,
	}
}

func (e *EmptyInputError) Error() string {
	return fmt.Sprintf("input is empty: %s", e.Value)
}

type StorageError struct {
	Msg string
}

func NewStorageError(msg string) *StorageError {
	return &StorageError{
		Msg: msg,
	}
}

func (e *StorageError) Error() string {
	return fmt.Sprintf("storage error: %s", e.Msg)
}

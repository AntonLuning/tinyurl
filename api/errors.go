package api

import "fmt"

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidationError(field string, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

type NotFoundError struct {
	Resource string `json:"resource"`
	Value    string `json:"value"`
}

func NewNotFoundError(resource string, value string) *NotFoundError {
	return &NotFoundError{
		Resource: resource,
		Value:    value,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found error for resource %s: %s", e.Resource, e.Value)
}

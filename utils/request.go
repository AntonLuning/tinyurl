package utils

import (
	"context"

	"github.com/google/uuid"
)

func SetContextValues(ctx context.Context, serverType string) context.Context {
	ctx = context.WithValue(ctx, REQUEST_ID_KEY, uuid.New().String())
	ctx = context.WithValue(ctx, REQUEST_SERVER_TYPE_KEY, serverType)

	return ctx
}

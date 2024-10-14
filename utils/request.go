package utils

import (
	"context"
	"sync/atomic"
)

var (
	_requestCounter uint64 = 0 // TODO: Should be persistent (in case of a restart)
)

func SetContextValues(ctx context.Context, serverType string) context.Context {
	ctx = context.WithValue(ctx, REQUEST_ID_KEY, atomic.AddUint64(&_requestCounter, 1))
	ctx = context.WithValue(ctx, REQUEST_SERVER_TYPE_KEY, serverType)

	return ctx
}

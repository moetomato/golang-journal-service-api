package middlewares

import (
	"context"
	"sync"
)

type traceIDKey struct{}

var (
	traceID int = 1
	m       sync.Mutex
)

func newTraceID() int {
	var id int
	m.Lock()
	id = traceID
	traceID += 1
	m.Unlock()
	return id
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

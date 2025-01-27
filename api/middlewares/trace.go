package middlewares

import (
	"sync"
)

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

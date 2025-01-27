package middlewares

import (
	"log"
	"net/http"

	"github.com/moetomato/golang-journal-service-api/common"
)

/*
	logging request URI, method and reposnse status code.
	signature of logging functions should be in the form of :
		func SampleMiddleware(next http.Handler) http.Handler
	to wrap original handlers.

	resLoggingWriter is to obtain response status code from http.ResponseWriter.
*/

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(statusCode int) {
	rsw.code = statusCode
	rsw.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()
		log.Printf("TraceID : %d, URI : %s, Method : %s\n", traceID, req.RequestURI, req.Method)

		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)

		// to log response code
		rsw := NewResLoggingWriter(w)
		next.ServeHTTP(rsw, req)
		log.Printf("TraceID : %d, StatusCode: %d", traceID, rsw.code)
	})
}

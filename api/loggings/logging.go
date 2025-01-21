package loggings

import (
	"log"
	"net/http"
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
		log.Println(req.RequestURI, req.Method)

		// to log response code
		rsw := NewResLoggingWriter(w)
		next.ServeHTTP(rsw, req)
		log.Println("response code", rsw.code)
	})
}

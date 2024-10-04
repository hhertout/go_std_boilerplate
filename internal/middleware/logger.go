package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	StatusCode int
	Size       int
}

func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{w, http.StatusOK, 0}
}

// WriteHeader permet de capturer le statut de la réponse
func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Write permet de capturer la taille de la réponse
func (rw *ResponseWriterWrapper) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.Size += size
	return size, err
}

func Logger(next http.Handler, l *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrappedWriter := NewResponseWriterWrapper(w)

		next.ServeHTTP(wrappedWriter, r)

		if r.URL.Path == "/health" {
			return
		}

		l.Info(
			"request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("params", r.URL.RawQuery),
			zap.String("content_type", r.Header.Get("Content-Type")),
			zap.String("user_agent", r.Header.Get("User-Agent")),
			zap.Int("status_code", wrappedWriter.StatusCode),
			zap.Int("response_size", wrappedWriter.Size),
			zap.Duration("duration", time.Since(start)),
		)
	})
}

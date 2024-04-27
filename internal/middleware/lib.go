package middleware

import (
	"context"
	"net/http"
)

type ContextKey string

func ChainMiddleware(handler http.Handler, middlewares ...func(next http.Handler) http.Handler) http.Handler {
	if len(middlewares) < 1 {
		return handler
	}

	var next http.Handler
	for i := 0; i < len(middlewares); i++ {
		next = middlewares[i](handler)
	}

	return next
}

func SetContext(r *http.Request, key string, value interface{}) *http.Request {
	var k = ContextKey(key)

	ctx := context.WithValue(r.Context(), k, value)

	return r.WithContext(ctx)
}

func GetContext(r *http.Request, key string) interface{} {
	var k = ContextKey(key)

	return r.Context().Value(k)
}

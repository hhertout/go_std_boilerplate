package middleware

import (
	"log"
	"net/http"
	"os"
)

func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("API_KEY") == "" {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal("API_KEY is not set")
		}

		apiKey := r.Header.Get("X-API-KEY")

		if apiKey != os.Getenv("API_KEY") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

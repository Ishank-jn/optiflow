package middleware

import (
	"log"
	"net/http"
	"optiflow/internal/oauth"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := tokenParts[1]

		// Validate token
		valid, err := oauth.ValidateToken(token)
		if err != nil || !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

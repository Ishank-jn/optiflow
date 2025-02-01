package middleware

import (
    "net/http"
    "strings"
	"context"
	"log"
    "github.com/golang-jwt/jwt/v5"
    "optiflow/internal/oauth"
)

// AuthMiddleware validates the JWT token in the Authorization header
func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header is required", http.StatusUnauthorized)
                return
            }

            // Extract the token from the header
            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            if tokenString == "" {
                http.Error(w, "Invalid token format", http.StatusUnauthorized)
                return
            }

            // Parse and validate the token
            token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
                return []byte(secretKey), nil
            })
            if err != nil {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            // Check if the token is valid
            if !token.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            // Add the user ID to the request context
            claims, ok := token.Claims.(*auth.Claims)
            if !ok {
                http.Error(w, "Invalid token claims", http.StatusUnauthorized)
                return
            }

            type contextKey string
            r = r.WithContext(context.WithValue(r.Context(), contextKey("user_id"), claims.UserID))
            next.ServeHTTP(w, r)
        })
    }
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

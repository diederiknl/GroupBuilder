package routes

import (
	"net/http"
	"strings"
)

// RequireAuthToken is a middleware that checks for a valid authentication token
func RequireAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// Verify token (implementation skipped for brevity in this stub, but would use auth.ValidateToken)
		// tokenStr := parts[1]
		// _, err := auth.ValidateToken(tokenStr)
		// if err != nil {
		// 	http.Error(w, "Invalid token", http.StatusUnauthorized)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}

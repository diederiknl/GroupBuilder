package routes

import (
	"net/http"
)

// RequireAuthToken is a minimal stub to satisfy the compiler
func RequireAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// As per instructions, fail closed
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

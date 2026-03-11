package routes

import (
	"net/http"
)

func RequireAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pass-through stub
		next.ServeHTTP(w, r)
	})
}

package routes

import (
	"context"
	"net/http"
	"strings"

	"GroupBuilder/internal/auth"
	"GroupBuilder/internal/database"
	"GroupBuilder/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type contextKey string

const ClaimsContextKey contextKey = "claims"

func SetupRoutes(db *database.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Basic routes
	r.Get("/", handlers.Welcome)

	// Authentication routes
	r.Post("/auth/student/login-link", handlers.SendLoginLink(db))
	r.Post("/auth/student/verify", handlers.VerifyStudentLoginLink(db))
	r.Post("/auth/teacher/login", handlers.TeacherLogin(db))

	// Student routes (protected)
	r.Group(func(r chi.Router) {
		r.Use(RequireAuthToken)
		r.Route("/students", func(r chi.Router) {
			r.Get("/", handlers.GetAllStudents(db))
			r.Post("/", handlers.CreateStudent(db))
			r.Get("/{id}", handlers.GetStudent(db))
			r.Put("/{id}", handlers.UpdateStudent(db))
			r.Delete("/{id}", handlers.DeleteStudent(db))
		})
	})

	// In de SetupRoutes functie, voeg deze regel toe binnen de groep die RequireTeacherRole gebruikt:
	r.Post("/import-students", handlers.ImportStudentList(db))
	// Add other routes here...

	return r
}

// RequireAuthToken is a middleware that verifies the JWT token
func RequireAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims to context
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

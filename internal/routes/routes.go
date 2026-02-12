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

const ClaimsKey contextKey = "claims"

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

		tokenStr := parts[1]
		claims, err := auth.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetupRoutes(db *database.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Basic routes
	r.Get("/", handlers.Welcome)

	// Authentication routes
	r.Post("/auth/student/login-link", handlers.SendLoginLink(db))
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

	// Other routes (placeholders if handlers are missing)
	// r.Post("/auth/student/verify", handlers.VerifyStudentLoginLink(db))
	// r.Post("/import-students", handlers.ImportStudentList(db))

	return r
}

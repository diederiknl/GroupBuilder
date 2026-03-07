package routes

import (
	"GroupBuilder/internal/database"
	"GroupBuilder/internal/handlers"
	"GroupBuilder/internal/auth"
	"strings"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

	// Teacher routes (protected)
	r.Group(func(r chi.Router) {
		r.Use(RequireTeacherRole)
		r.Post("/import-students", handlers.ImportStudentList(db))
	})

	return r
}

func RequireTeacherRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: missing or invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateToken(tokenStr)
		if err != nil || claims == nil || claims.Role != "teacher" {
			http.Error(w, "Forbidden: Teacher role required", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RequireAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: missing or invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateToken(tokenStr)
		if err != nil || claims == nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

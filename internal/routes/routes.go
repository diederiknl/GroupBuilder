package routes

import (
	"github.com/GroupBuilder/internal/database"
	"github.com/GroupBuilder/internal/handlers"
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

	// In de SetupRoutes functie, voeg deze regel toe binnen de groep die RequireTeacherRole gebruikt:
	r.Post("/import-students", handlers.ImportStudentList(db))
	// Add other routes here...

	return r
}

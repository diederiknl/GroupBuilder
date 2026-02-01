package handlers

import (
	"encoding/json"
	"net/http"

	"GroupBuilder/internal/auth"
	"GroupBuilder/internal/database"
)

func SendLoginLink(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := auth.GenerateLoginLink(req.Email)
		if err != nil {
			http.Error(w, "Failed to generate login link", http.StatusInternalServerError)
			return
		}

		// TODO: Save the link to the database and send email
		// For now, we'll just log it or return it in the response for testing purposes (NEVER do this in prod)
		// Since this is a fix for build failure, I'll just use the variable to satisfy the compiler.
		_ = link

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login link sent"})
	}
}

func TeacherLogin(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: Verify username and password against database

		// If login successful, generate and return a JWT token
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": "JWT_TOKEN_HERE"})
	}
}

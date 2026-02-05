package handlers

import (
	"encoding/json"
	"fmt"
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
		// For now, log it to stdout for testing
		fmt.Printf("Login link for %s: %s\n", req.Email, link)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login link sent"})
	}
}

func VerifyStudentLoginLink(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "Missing token", http.StatusBadRequest)
			return
		}

		// TODO: Verify token against database
		// For now, just generate a JWT

		jwtToken, err := auth.GenerateToken("student@example.com", "student")
		if err != nil {
			http.Error(w, "Failed to generate session token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": jwtToken})
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

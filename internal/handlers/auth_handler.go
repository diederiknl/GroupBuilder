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
		_ = link // Prevent unused variable error while email sending is unimplemented

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

func VerifyStudentLoginLink(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "Token required", http.StatusBadRequest)
			return
		}

		claims, err := auth.ValidateToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if it's a student token?
		if claims.Role != "student" {
			http.Error(w, "Invalid role", http.StatusForbidden)
			return
		}

		// Exchange for a session token or just return success?
		// For now, just return valid.
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Login successful",
			"email":   claims.Email,
			"token":   token, // Reuse the token or issue a new one
		})
	}
}

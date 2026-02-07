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
		// For now, we'll return it in the response for testing purposes (in a real app, don't do this!)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login link sent",
			"link":    link,
		})
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
		tokenStr := r.URL.Query().Get("token")
		if tokenStr == "" {
			http.Error(w, "Missing token", http.StatusBadRequest)
			return
		}

		claims, err := auth.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// In a real app, you might want to exchange this for a session token or set a cookie
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Login successful",
			"email":   claims.Email,
			"role":    claims.Role,
		})
	}
}

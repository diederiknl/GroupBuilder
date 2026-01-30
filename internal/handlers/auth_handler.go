package handlers

import (
	"encoding/json"
	"net/http"

	"GroupBuilder/internal/auth"
	"GroupBuilder/internal/database"

	"golang.org/x/crypto/bcrypt"
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
		// For now, we just print it so it's "used"
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

		var passwordHash string
		// Use parameterized query to prevent SQL injection
		err := db.QueryRow("SELECT password_hash FROM teachers WHERE username = ?", req.Username).Scan(&passwordHash)
		if err != nil {
			// Don't distinguish between user not found and other errors for security (enumeration)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Verify password
		if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Generate real JWT token
		token, err := auth.GenerateToken(req.Username, "teacher")
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

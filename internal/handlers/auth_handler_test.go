package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"GroupBuilder/internal/database"

	"golang.org/x/crypto/bcrypt"
)

func setupTestDB(t *testing.T) *database.DB {
	db, err := database.InitDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to init db: %v", err)
	}
	return db
}

func TestTeacherLogin(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Create a teacher
	password := "securepassword"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := db.Exec("INSERT INTO teachers (username, password_hash) VALUES (?, ?)", "teacher1", string(hash))
	if err != nil {
		t.Fatalf("Failed to insert teacher: %v", err)
	}

	tests := []struct {
		name           string
		username       string
		password       string
		expectedStatus int
	}{
		{
			name:           "Valid Credentials",
			username:       "teacher1",
			password:       "securepassword",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Password",
			username:       "teacher1",
			password:       "wrongpassword",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Non-existent User",
			username:       "teacher2",
			password:       "anypassword",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	handler := TeacherLogin(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := map[string]string{
				"username": tt.username,
				"password": tt.password,
			}
			body, _ := json.Marshal(payload)
			req, _ := http.NewRequest("POST", "/auth/teacher/login", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}

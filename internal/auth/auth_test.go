package auth

import (
	"os"
	"strings"
	"testing"
)

func TestGetJWTKey(t *testing.T) {
	// Save the original value and restore it after the test
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Test missing env variable
	os.Setenv("JWT_SECRET", "")
	_, err := getJWTKey()
	if err == nil {
		t.Errorf("Expected error for missing JWT_SECRET, got nil")
	}

	// Test valid env variable
	os.Setenv("JWT_SECRET", "test_secret")
	key, err := getJWTKey()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(key) != "test_secret" {
		t.Errorf("Expected key 'test_secret', got '%s'", string(key))
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "test_secret_for_token")

	email := "test@example.com"
	role := "student"

	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email '%s', got '%s'", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role '%s', got '%s'", role, claims.Role)
	}
}

func TestGenerateLoginLink(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "test_secret_for_link")

	email := "link@example.com"
	link, err := GenerateLoginLink(email)
	if err != nil {
		t.Fatalf("Failed to generate login link: %v", err)
	}

	if !strings.HasPrefix(link, "/login?token=") {
		t.Errorf("Expected link to start with '/login?token=', got '%s'", link)
	}

	// Validate the token in the link
	tokenStr := strings.TrimPrefix(link, "/login?token=")
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token from link: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email '%s', got '%s'", email, claims.Email)
	}
	if claims.Role != "student" {
		t.Errorf("Expected role 'student', got '%s'", claims.Role)
	}
}

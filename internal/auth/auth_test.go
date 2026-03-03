package auth

import (
	"os"
	"testing"
)

func TestInitKeyAndToken(t *testing.T) {
	// Save existing environment variable to restore later
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Test 1: Set a new secret and test token generation
	testSecret := "super_secret_test_key_123"
	err := os.Setenv("JWT_SECRET", testSecret)
	if err != nil {
		t.Fatalf("Failed to set env var: %v", err)
	}

	InitKey()

	if string(jwtKey) != testSecret {
		t.Errorf("InitKey() didn't set jwtKey correctly. Expected %v, got %v", testSecret, string(jwtKey))
	}

	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

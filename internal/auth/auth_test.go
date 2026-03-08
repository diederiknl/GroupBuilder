package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Temporarily set JWT_SECRET for the test
	originalSecret := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer func() {
		if originalSecret == "" {
			os.Unsetenv("JWT_SECRET")
		} else {
			os.Setenv("JWT_SECRET", originalSecret)
		}
	}()

	email := "test@example.com"
	role := "student"

	// Test GenerateToken
	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if tokenStr == "" {
		t.Fatalf("Generated token is empty")
	}

	// Test ValidateToken
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestNoJWTSecret(t *testing.T) {
	// Temporarily unset JWT_SECRET for the test
	originalSecret := os.Getenv("JWT_SECRET")
	os.Unsetenv("JWT_SECRET")
	defer func() {
		if originalSecret != "" {
			os.Setenv("JWT_SECRET", originalSecret)
		}
	}()

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatalf("Expected error when JWT_SECRET is missing, got nil")
	}

	_, err = ValidateToken("some_token_string")
	if err == nil {
		t.Fatalf("Expected error when JWT_SECRET is missing, got nil")
	}
}

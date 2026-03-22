package auth

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Save the original value
	origSecret := os.Getenv("JWT_SECRET")
	// Restore the original value when the test completes
	defer os.Setenv("JWT_SECRET", origSecret)

	// Set test environment variable
	os.Setenv("JWT_SECRET", "test_secret_key")

	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Fatalf("Expected token string, got empty string")
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

func TestGenerateTokenWithoutSecret(t *testing.T) {
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	// Unset environment variable
	os.Setenv("JWT_SECRET", "")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatalf("Expected error when JWT_SECRET is not set")
	}
	if !strings.Contains(err.Error(), "environment variable is not set") {
		t.Errorf("Expected error mentioning environment variable, got: %v", err)
	}
}

func TestValidateTokenWithoutSecret(t *testing.T) {
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	// First set it to generate a valid token
	os.Setenv("JWT_SECRET", "test_secret_key")
	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Then unset it before validation
	os.Setenv("JWT_SECRET", "")
	_, err = ValidateToken(token)
	if err == nil {
		t.Fatalf("Expected error when JWT_SECRET is not set")
	}
	if !strings.Contains(err.Error(), "environment variable is not set") {
		t.Errorf("Expected error mentioning environment variable, got: %v", err)
	}
}

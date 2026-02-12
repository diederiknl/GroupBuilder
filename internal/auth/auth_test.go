package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set the environment variable for testing
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(token)
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

func TestMissingJwtSecret(t *testing.T) {
	// Ensure the environment variable is not set
	os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	_, err := GenerateToken(email, role)
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}

	// Also check ValidateToken
	_, err = ValidateToken("some.token.here")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

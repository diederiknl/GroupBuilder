package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set the environment variable for testing
	originalSecret := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Setenv("JWT_SECRET", originalSecret)

	email := "test@example.com"
	role := "student"

	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestGenerateTokenWithoutSecret(t *testing.T) {
	// Ensure the environment variable is NOT set
	originalSecret := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "")
	defer os.Setenv("JWT_SECRET", originalSecret)

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is not set, got nil")
	}
	if err.Error() != "JWT_SECRET environment variable is not set" {
		t.Errorf("Expected specific error message, got: %v", err)
	}
}

func TestValidateTokenWithoutSecret(t *testing.T) {
	// First generate a token with a valid secret
	originalSecret := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "test_secret_key")
	tokenStr, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Expected no error generating token, got %v", err)
	}

	// Now try to validate it with the secret removed
	os.Setenv("JWT_SECRET", "")
	defer os.Setenv("JWT_SECRET", originalSecret)

	_, err = ValidateToken(tokenStr)
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is not set during validation, got nil")
	}
}

package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set the environment variable for testing
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Setenv("JWT_SECRET", "") // cleanup

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

func TestGenerateTokenMissingSecret(t *testing.T) {
	// Ensure the environment variable is not set
	os.Setenv("JWT_SECRET", "")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is not set, got nil")
	}
}

func TestValidateTokenMissingSecret(t *testing.T) {
	// Ensure the environment variable is not set
	os.Setenv("JWT_SECRET", "")

	_, err := ValidateToken("dummy.token.string")
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is not set, got nil")
	}
}

func TestValidateTokenInvalid(t *testing.T) {
	// Set the environment variable for testing
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Setenv("JWT_SECRET", "") // cleanup

	_, err := ValidateToken("invalid.token.string")
	if err == nil {
		t.Fatal("Expected error when token is invalid, got nil")
	}
}

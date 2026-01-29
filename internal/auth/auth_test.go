package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
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

func TestEnvVarKey(t *testing.T) {
	// Save original env
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Set a custom secret
	customSecret := "my_custom_secret_key_12345"
	os.Setenv("JWT_SECRET", customSecret)

	email := "test@example.com"
	role := "student"

	// Generate token with custom secret
	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token with custom secret: %v", err)
	}

	// Validate with same custom secret (should pass)
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token with custom secret: %v", err)
	}
	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	// Change secret to something else
	os.Setenv("JWT_SECRET", "different_secret")

	// Validate should fail
	_, err = ValidateToken(token)
	if err == nil {
		t.Error("Expected validation to fail with different secret, but it passed")
	}
}

package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Ensure we are using the default key
	os.Unsetenv("JWT_SECRET")

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
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestTokenWithEnvVar(t *testing.T) {
	secret := "my-super-secret-key-123"
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")

	email := "secure@example.com"
	role := "admin"

	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token with custom secret: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
}

func TestTokenKeyMismatch(t *testing.T) {
	// Generate with default key
	os.Unsetenv("JWT_SECRET")
	email := "mismatch@example.com"
	role := "student"

	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Try to validate with custom key
	os.Setenv("JWT_SECRET", "different-secret")
	defer os.Unsetenv("JWT_SECRET")

	_, err = ValidateToken(tokenStr)
	if err == nil {
		t.Error("Expected validation error due to key mismatch, got nil")
	}
}

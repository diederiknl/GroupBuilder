package auth

import (
	"os"
	"testing"
)

func TestTokenGeneration(t *testing.T) {
	// Set secret
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET")

	// Generate a token
	email := "test@example.com"
	role := "student"
	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Verify the token works
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}
	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
}

func TestMissingSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("foo", "bar")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

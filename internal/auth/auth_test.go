package auth

import (
	"os"
	"testing"
)

func TestTokenFlow(t *testing.T) {
	// Setup env
	secret := "test_secret_123"
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")

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

func TestMissingSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")
	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

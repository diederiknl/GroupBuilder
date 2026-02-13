package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken_Success(t *testing.T) {
	original := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", original)

	os.Setenv("JWT_SECRET", "test_secret")

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

func TestGenerateToken_MissingSecret(t *testing.T) {
	original := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", original)
	os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	_, err := GenerateToken(email, role)
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Save the original value and restore it at the end of the test.
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	testSecret := "supersecretkey"
	os.Setenv("JWT_SECRET", testSecret)

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

func TestGenerateToken_NoSecret(t *testing.T) {
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	os.Setenv("JWT_SECRET", "")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatalf("Expected error when JWT_SECRET is empty, got nil")
	}
}

func TestValidateToken_NoSecret(t *testing.T) {
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	// First generate a token with a valid secret
	testSecret := "supersecretkey"
	os.Setenv("JWT_SECRET", testSecret)
	tokenStr, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Unset secret to test validation failure
	os.Setenv("JWT_SECRET", "")

	_, err = ValidateToken(tokenStr)
	if err == nil {
		t.Fatalf("Expected error when JWT_SECRET is empty, got nil")
	}
}

package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set JWT_SECRET for this test
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "teacher"

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

func TestMissingJwtSecret(t *testing.T) {
	// Ensure JWT_SECRET is not set
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "teacher")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

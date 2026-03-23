package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Setup test environment
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	os.Setenv("JWT_SECRET", "test_secret_key")

	email := "test@example.com"
	role := "student"

	// Test GenerateToken
	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if tokenStr == "" {
		t.Fatalf("GenerateToken returned empty token")
	}

	// Test ValidateToken
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}
	if claims == nil {
		t.Fatalf("ValidateToken returned nil claims")
	}
	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestGenerateToken_NoSecret(t *testing.T) {
	// Setup test environment
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	os.Setenv("JWT_SECRET", "") // Unset the secret

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is not set, got nil")
	}
}

func TestValidateToken_NoSecret(t *testing.T) {
	// Setup test environment
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Generate a valid token first
	os.Setenv("JWT_SECRET", "test_secret_key")
	tokenStr, _ := GenerateToken("test@example.com", "student")

	// Unset the secret for validation
	os.Setenv("JWT_SECRET", "")

	_, err := ValidateToken(tokenStr)
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is not set, got nil")
	}
}

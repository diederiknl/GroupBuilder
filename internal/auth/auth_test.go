package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Setup test environment variable
	originalSecret := os.Getenv("JWT_SECRET")
	err := os.Setenv("JWT_SECRET", "test_secret_key")
	if err != nil {
		t.Fatalf("Failed to set JWT_SECRET: %v", err)
	}
	defer os.Setenv("JWT_SECRET", originalSecret)

	email := "test@example.com"
	role := "student"

	// Test token generation
	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if tokenStr == "" {
		t.Fatal("GenerateToken returned empty string")
	}

	// Test token validation
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}
	if claims == nil {
		t.Fatal("ValidateToken returned nil claims")
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestValidateTokenInvalid(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Setenv("JWT_SECRET", originalSecret)

	invalidToken := "invalid.token.string"
	claims, err := ValidateToken(invalidToken)
	if err == nil {
		t.Fatal("ValidateToken expected error for invalid token, got none")
	}
	if claims != nil {
		t.Fatal("ValidateToken expected nil claims for invalid token")
	}
}

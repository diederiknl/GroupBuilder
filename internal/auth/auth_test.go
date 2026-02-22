package auth

import (
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	// specific test setup: ensure clean state
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Ensure JWT_SECRET is unset
	os.Unsetenv("JWT_SECRET")

	// Case 1: Missing Secret
	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is not set, got nil")
	}

	// Case 2: Valid Secret
	os.Setenv("JWT_SECRET", "testsecret")
	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Errorf("Expected no error when JWT_SECRET is set, got %v", err)
	}
	if token == "" {
		t.Error("Expected token to be generated, got empty string")
	}
}

func TestValidateToken(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	os.Setenv("JWT_SECRET", "testsecret")

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Errorf("Expected valid token, got error: %v", err)
	}
	if claims.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got %s", claims.Email)
	}
}

func TestValidateTokenInvalidSecret(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Generate token with one secret
	os.Setenv("JWT_SECRET", "secret1")
	token, _ := GenerateToken("test@example.com", "student")

	// Validate with another secret
	os.Setenv("JWT_SECRET", "secret2")
	_, err := ValidateToken(token)
	if err == nil {
		t.Error("Expected error when validating token with different secret, got nil")
	}
}

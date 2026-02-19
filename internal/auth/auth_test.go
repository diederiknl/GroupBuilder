package auth

import (
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	// Set environment variable for testing
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET")

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	if token == "" {
		t.Fatal("Token is empty")
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", claims.Email)
	}
	if claims.Role != "student" {
		t.Errorf("Expected role student, got %s", claims.Role)
	}
}

func TestGenerateToken_MissingSecret(t *testing.T) {
	// Ensure environment variable is NOT set
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is missing, got nil")
	}
}

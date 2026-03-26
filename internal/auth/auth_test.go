package auth

import (
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	// Save the original JWT_SECRET
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Test case 1: JWT_SECRET is missing
	os.Setenv("JWT_SECRET", "")
	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected an error when JWT_SECRET is not set, but got none")
	}

	// Test case 2: JWT_SECRET is set
	os.Setenv("JWT_SECRET", "test_secret")
	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Errorf("Unexpected error when generating token: %v", err)
	}
	if token == "" {
		t.Error("Generated token should not be empty")
	}
}

func TestValidateToken(t *testing.T) {
	// Save the original JWT_SECRET
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Test case 1: JWT_SECRET is not set for validation
	os.Setenv("JWT_SECRET", "")
	_, err := ValidateToken("some.fake.token")
	if err == nil {
		t.Error("Expected an error when validating with missing JWT_SECRET, but got none")
	}

	// Test case 2: JWT_SECRET is set
	os.Setenv("JWT_SECRET", "test_secret")
	tokenStr, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token for test setup: %v", err)
	}

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Errorf("Failed to validate token: %v", err)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", claims.Email)
	}
	if claims.Role != "student" {
		t.Errorf("Expected role 'student', got '%s'", claims.Role)
	}
}

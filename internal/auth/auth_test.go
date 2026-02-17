package auth

import (
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	// Set JWT_SECRET for test
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Fatal("Expected token to be generated")
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Expected no error validating token, got %v", err)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %v", claims.Email)
	}
	if claims.Role != "student" {
		t.Errorf("Expected role student, got %v", claims.Role)
	}
}

func TestMissingSecret(t *testing.T) {
	// Unset JWT_SECRET just in case
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

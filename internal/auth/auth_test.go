package auth

import (
	"os"
	"testing"
)

func TestTokenGenerationAndValidation(t *testing.T) {
	// Set the environment variable for testing
	err := os.Setenv("JWT_SECRET", "test_secret_key")
	if err != nil {
		t.Fatalf("Failed to set JWT_SECRET: %v", err)
	}
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	// Test generation
	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if tokenStr == "" {
		t.Fatal("Expected a token, got empty string")
	}

	// Test validation
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestGenerateLoginLink(t *testing.T) {
	// Set the environment variable for testing
	err := os.Setenv("JWT_SECRET", "test_secret_key")
	if err != nil {
		t.Fatalf("Failed to set JWT_SECRET: %v", err)
	}
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	link, err := GenerateLoginLink(email)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Checking if it starts with the correct path
	if len(link) < 13 || link[:13] != "/login?token=" {
		t.Errorf("Expected link to start with /login?token=, got %s", link)
	}
}

func TestMissingJWTSecret(t *testing.T) {
	// Ensure the environment variable is not set
	os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	// Test GenerateToken
	_, err := GenerateToken(email, role)
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}

	// Test ValidateToken
	_, err = ValidateToken("dummy.token.string")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}

	// Test GenerateLoginLink
	_, err = GenerateLoginLink(email)
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

package auth

import (
	"os"
	"strings"
	"testing"
)

func TestTokenGenerationAndValidation(t *testing.T) {
	// Backup original JWT_SECRET
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	// Set a test secret
	testSecret := "test_secret_key"
	os.Setenv("JWT_SECRET", testSecret)

	email := "test@example.com"
	role := "student"

	// Test GenerateToken
	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	if tokenStr == "" {
		t.Fatal("Expected non-empty token string")
	}

	// Test ValidateToken
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

func TestGenerateLoginLink(t *testing.T) {
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	os.Setenv("JWT_SECRET", "test_secret_key")

	email := "user@example.com"
	link, err := GenerateLoginLink(email)
	if err != nil {
		t.Fatalf("Failed to generate login link: %v", err)
	}

	if !strings.HasPrefix(link, "/login?token=") {
		t.Errorf("Expected link to start with '/login?token=', got %s", link)
	}
}

func TestMissingSecretReturnsError(t *testing.T) {
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	// Clear the secret
	os.Setenv("JWT_SECRET", "")

	// Test GenerateToken error
	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is not set, got nil")
	}
	if err.Error() != "JWT_SECRET environment variable is not set" {
		t.Errorf("Unexpected error message: %v", err)
	}

	// Test ValidateToken error
	_, err = ValidateToken("some.token.string")
	if err == nil {
		t.Fatal("Expected error when validating without JWT_SECRET, got nil")
	}

	// Test GenerateLoginLink error
	_, err = GenerateLoginLink("test@example.com")
	if err == nil {
		t.Fatal("Expected error when generating login link without JWT_SECRET, got nil")
	}
}

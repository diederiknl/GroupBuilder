package auth

import (
	"os"
	"strings"
	"testing"
)

func TestTokenGenerationAndValidation(t *testing.T) {
	// Setup
	originalKey := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalKey)
	os.Setenv("JWT_SECRET", "test_secret_key")

	email := "student@example.com"
	role := "student"

	// Test GenerateToken
	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("GenerateToken returned empty token")
	}

	// Test ValidateToken
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

	// Test GenerateLoginLink
	link, err := GenerateLoginLink(email)
	if err != nil {
		t.Fatalf("GenerateLoginLink failed: %v", err)
	}
	if !strings.HasPrefix(link, "/login?token=") {
		t.Errorf("Invalid link format: %s", link)
	}
}

func TestMissingJWTSecret(t *testing.T) {
	// Setup
	originalKey := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalKey)
	os.Setenv("JWT_SECRET", "")

	email := "student@example.com"
	role := "student"

	// Test GenerateToken fails
	_, err := GenerateToken(email, role)
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is missing, got none")
	}

	// Test ValidateToken fails
	_, err = ValidateToken("dummy_token")
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is missing, got none")
	}
}

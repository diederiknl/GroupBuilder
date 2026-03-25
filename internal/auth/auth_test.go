package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Save the original value and restore it later
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
		t.Fatal("Expected a token, got an empty string")
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

func TestGenerateToken_MissingSecret(t *testing.T) {
	// Save the original value and restore it later
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	// Clear the secret
	os.Setenv("JWT_SECRET", "")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is missing, got nil")
	}
	expectedErr := "JWT_SECRET environment variable is not set"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
	}
}

func TestValidateToken_MissingSecret(t *testing.T) {
	// Save the original value and restore it later
	origSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", origSecret)

	// Set a secret to generate a token first
	os.Setenv("JWT_SECRET", "test_secret")
	tokenStr, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token for setup: %v", err)
	}

	// Clear the secret for testing validation
	os.Setenv("JWT_SECRET", "")

	_, err = ValidateToken(tokenStr)
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is missing, got nil")
	}
	expectedErr := "JWT_SECRET environment variable is not set"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
	}
}

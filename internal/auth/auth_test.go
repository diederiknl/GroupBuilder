package auth

import (
	"os"
	"testing"
)

func TestInitKey(t *testing.T) {
	// Clear the env var initially
	os.Unsetenv("JWT_SECRET")

	// Test case 1: ENV var not set
	err := InitKey()
	if err == nil {
		t.Errorf("expected error when JWT_SECRET is not set, got nil")
	}

	// Test case 2: ENV var set
	expectedKey := "test_secret_key"
	os.Setenv("JWT_SECRET", expectedKey)
	defer os.Unsetenv("JWT_SECRET") // Clean up

	err = InitKey()
	if err != nil {
		t.Errorf("expected no error when JWT_SECRET is set, got %v", err)
	}

	if string(jwtKey) != expectedKey {
		t.Errorf("expected jwtKey to be %q, got %q", expectedKey, string(jwtKey))
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	// Setup the jwtKey
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET")
	InitKey()

	email := "test@example.com"
	role := "student"

	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}
	if tokenStr == "" {
		t.Errorf("expected non-empty token string")
	}

	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("failed to validate token: %v", err)
	}
	if claims.Email != email {
		t.Errorf("expected email %q, got %q", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("expected role %q, got %q", role, claims.Role)
	}
}

package auth

import (
	"os"
	"strings"
	"testing"
)

func TestAuthWithSecret(t *testing.T) {
	// Save the old value of JWT_SECRET to restore it later
	oldSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", oldSecret)

	os.Setenv("JWT_SECRET", "test_secret_key")

	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
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
		t.Fatalf("Failed to generate login link: %v", err)
	}

	if !strings.HasPrefix(link, "/login?token=") {
		t.Errorf("Generated link does not have correct prefix: %s", link)
	}
}

func TestAuthWithoutSecret(t *testing.T) {
	// Save the old value of JWT_SECRET to restore it later
	oldSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", oldSecret)

	os.Setenv("JWT_SECRET", "")

	email := "test@example.com"
	role := "student"

	_, err := GenerateToken(email, role)
	if err == nil {
		t.Errorf("Expected error when generating token without JWT_SECRET, got nil")
	}

	// We can't really test ValidateToken without a token, but let's test a dummy one
	_, err = ValidateToken("dummy_token")
	if err == nil {
		t.Errorf("Expected error when validating token without JWT_SECRET, got nil")
	}
}

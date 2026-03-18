package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Need to use os.Setenv because Go 1.16 doesn't have t.Setenv
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	if token == "" {
		t.Fatal("Expected a token, got empty string")
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
}

func TestGenerateTokenWithoutSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is not set, got nil")
	}
}

func TestValidateTokenWithoutSecret(t *testing.T) {
	// First generate a valid token
	os.Setenv("JWT_SECRET", "test_secret")
	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Now try to validate it without the secret
	os.Unsetenv("JWT_SECRET")
	_, err = ValidateToken(token)
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is not set, got nil")
	}
}

func TestGenerateLoginLink(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	link, err := GenerateLoginLink(email)
	if err != nil {
		t.Fatalf("Failed to generate login link: %v", err)
	}

	expectedPrefix := "/login?token="
	if len(link) <= len(expectedPrefix) || link[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("Expected link to start with %s, got %s", expectedPrefix, link)
	}

	// Extract and validate token
	tokenStr := link[len(expectedPrefix):]
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("Failed to validate token from link: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
}

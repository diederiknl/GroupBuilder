package auth

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set a specific secret for testing
	testSecret := "test_secret_123"
	os.Setenv("JWT_SECRET", testSecret)
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	// Generate token
	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Fatal("Token is empty")
	}

	// Validate token
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

func TestTokenValidationWithWrongSecret(t *testing.T) {
	// 1. Generate token with Secret A
	os.Setenv("JWT_SECRET", "secret_A")
	token, err := GenerateToken("user@example.com", "user")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// 2. Try to validate with Secret B
	os.Setenv("JWT_SECRET", "secret_B")
	defer os.Unsetenv("JWT_SECRET")

	_, err = ValidateToken(token)
	if err == nil {
		t.Error("Expected validation error when secrets don't match, but got nil")
	}
}

func TestGenerateLoginLink(t *testing.T) {
	os.Setenv("JWT_SECRET", "link_secret")
	defer os.Unsetenv("JWT_SECRET")

	// Test with custom BASE_URL
	customBase := "https://myapp.com"
	os.Setenv("BASE_URL", customBase)
	defer os.Unsetenv("BASE_URL")

	email := "student@example.com"
	link, err := GenerateLoginLink(email)
	if err != nil {
		t.Fatalf("Failed to generate login link: %v", err)
	}

	expectedPrefix := customBase + "/auth/student/verify?token="
	if !strings.Contains(link, expectedPrefix) {
		t.Errorf("Link format is incorrect. Expected prefix %s, got %s", expectedPrefix, link)
	}
}

func TestDefaultBaseURL(t *testing.T) {
	os.Unsetenv("BASE_URL")

	link, err := GenerateLoginLink("test@example.com")
	if err != nil {
		t.Fatalf("Failed to generate login link: %v", err)
	}

	if !strings.Contains(link, "http://localhost:8080") {
		t.Errorf("Expected default localhost URL, got %s", link)
	}
}

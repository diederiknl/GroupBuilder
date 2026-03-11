package auth

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateAndValidateToken_Success(t *testing.T) {
	// Setup
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	// Generate
	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Fatal("Expected a token, got empty string")
	}

	// Validate
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Expected no error on validation, got %v", err)
	}
	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestGenerateToken_NoSecret(t *testing.T) {
	// Ensure unset
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is missing, got none")
	}
	if !strings.Contains(err.Error(), "JWT_SECRET") {
		t.Errorf("Expected error to mention JWT_SECRET, got %v", err)
	}
}

func TestValidateToken_NoSecret(t *testing.T) {
	// Temporarily set to generate a valid token
	os.Setenv("JWT_SECRET", "test_secret")
	token, _ := GenerateToken("test@example.com", "student")

	// Now unset to test validation failure
	os.Unsetenv("JWT_SECRET")

	_, err := ValidateToken(token)
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is missing, got none")
	}
	if !strings.Contains(err.Error(), "JWT_SECRET") {
		t.Errorf("Expected error to mention JWT_SECRET, got %v", err)
	}
}

func TestGenerateLoginLink_Success(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	link, err := GenerateLoginLink("test@example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !strings.HasPrefix(link, "/auth/student/verify?token=") {
		t.Errorf("Unexpected link format: %s", link)
	}
}

func TestGenerateLoginLink_NoSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateLoginLink("test@example.com")
	if err == nil {
		t.Fatal("Expected an error when JWT_SECRET is missing, got none")
	}
}

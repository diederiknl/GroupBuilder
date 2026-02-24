package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set JWT_SECRET for the test
	os.Setenv("JWT_SECRET", "test_secret_key")
	defer os.Unsetenv("JWT_SECRET")

	email := "test@example.com"
	role := "student"

	tokenStr, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

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

func TestValidateToken_InvalidSecret(t *testing.T) {
	// Set JWT_SECRET
	os.Setenv("JWT_SECRET", "correct_secret")
	defer os.Unsetenv("JWT_SECRET")

	// Generate token with correct secret
	tokenStr, _ := GenerateToken("user@example.com", "user")

	// Change secret
	os.Setenv("JWT_SECRET", "wrong_secret")
	// If the implementation actually reads from env, then validation should fail here.
	// However, if it uses hardcoded key, it will still pass, and this test will FAIL.

	// Validate should fail
	_, err := ValidateToken(tokenStr)
	if err == nil {
		t.Error("Expected validation error with changed secret, but got success")
	}
}

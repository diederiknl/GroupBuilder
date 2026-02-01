package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
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
}

func TestEnvVarChange(t *testing.T) {
	// Save original env var and restore after test
	originalKey := os.Getenv("JWT_SECRET")
	defer func() {
		if originalKey != "" {
			os.Setenv("JWT_SECRET", originalKey)
		} else {
			os.Unsetenv("JWT_SECRET")
		}
	}()

	email := "test@example.com"
	role := "student"

	// Case 1: Default key
	os.Unsetenv("JWT_SECRET")
	token1, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token 1: %v", err)
	}

	// Case 2: Custom key
	os.Setenv("JWT_SECRET", "new_secret_key_12345")
	token2, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token 2: %v", err)
	}

	if token1 == token2 {
		t.Error("Tokens generated with different keys should be different (collision or key ignored)")
	}

	// Verify token2 works with current env (which is set to custom key)
	_, err = ValidateToken(token2)
	if err != nil {
		t.Errorf("Failed to validate token2 with custom key: %v", err)
	}

	// Verify token1 fails with current env (custom key)
	_, err = ValidateToken(token1)
	if err == nil {
		t.Error("Token1 (default key) should not be valid when JWT_SECRET is set to custom key")
	}
}

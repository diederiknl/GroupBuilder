package auth

import (
	"os"
	"testing"
)

func TestTokenGenerationAndValidation(t *testing.T) {
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

func TestGetJWTKey(t *testing.T) {
	// Save original env
	original := os.Getenv("JWT_SECRET")
	defer func() {
		if original != "" {
			os.Setenv("JWT_SECRET", original)
		} else {
			os.Unsetenv("JWT_SECRET")
		}
	}()

	// Test default
	os.Unsetenv("JWT_SECRET")
	key := getJWTKey()
	if string(key) != "development_secret_key_change_me" {
		t.Errorf("Expected default key, got %s", key)
	}

	// Test with env var
	expected := "new_secret_key"
	os.Setenv("JWT_SECRET", expected)
	key = getJWTKey()
	if string(key) != expected {
		t.Errorf("Expected %s, got %s", expected, key)
	}
}

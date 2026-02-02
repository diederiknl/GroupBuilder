package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
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

func TestEnvironmentVariable(t *testing.T) {
	// 1. Generate token with default key
	os.Unsetenv("JWT_SECRET")
	tokenDefault, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token with default key: %v", err)
	}

	// 2. Set custom key
	customKey := "my_custom_secure_key"
	os.Setenv("JWT_SECRET", customKey)
	defer os.Unsetenv("JWT_SECRET") // Clean up

	// 3. Validate token with custom key (should fail)
	_, err = ValidateToken(tokenDefault)
	if err == nil {
		t.Error("Expected validation to fail when keys mismatch, but it succeeded")
	}

	// 4. Generate token with custom key
	tokenCustom, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token with custom key: %v", err)
	}

	// 5. Validate token with custom key (should pass)
	claims, err := ValidateToken(tokenCustom)
	if err != nil {
		t.Fatalf("Failed to validate token with custom key: %v", err)
	}
	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", claims.Email)
	}
}

package auth

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set JWT_SECRET
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
		t.Errorf("Expected email %v, got %v", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role %v, got %v", role, claims.Role)
	}
}

func TestTokenUsesEnvVar(t *testing.T) {
	// This test ensures that the token is signed using the environment variable,
	// not a hardcoded key.
	secret := "dynamic_secret_key"
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")

	tokenStr, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Manually parse with the expected secret
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// Before the fix, this will fail because it uses the hardcoded key
	if err != nil || !token.Valid {
		t.Logf("Token validation failed as expected (before fix): %v", err)
		// We expect failure here before the fix, but for the test to pass after the fix,
		// we should assert success. However, right now I want to confirm failure.
		// So I will leave it as an error to confirm the test fails initially.
		t.Errorf("Token was not signed with the environment variable secret.")
	}
}

func TestMissingEnvVar(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	// Before the fix, this succeeds (using hardcoded key)
	if err == nil {
		t.Error("Expected error when JWT_SECRET is not set, got nil")
	}
}

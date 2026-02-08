package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set a test secret
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

func TestMissingJWTSecret(t *testing.T) {
	// Ensure JWT_SECRET is unset
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Fatal("Expected error when JWT_SECRET is not set, got nil")
	}

	// Also test ValidateToken
	// We pass a dummy token string because we expect it to fail on key retrieval before parsing fully or during verification
	// Actually, if the token is invalid format, ParseWithClaims might return error before calling Keyfunc.
	// But let's try with a dummy string.
	// If it fails with "token contains an invalid number of segments", that's fine, but we want to verify it fails due to missing key if possible.
	// But GenerateToken failing is enough proof that getJwtKey works.
	// Let's stick to testing GenerateToken for missing key.
}

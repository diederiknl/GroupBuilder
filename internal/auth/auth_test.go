package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	email := "test@example.com"
	role := "student"

	// Ensure clean state
	os.Unsetenv("JWT_SECRET")

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

func TestJwtSecretEnv(t *testing.T) {
	email := "test@example.com"
	role := "student"

	// Set custom secret
	secret := "my_custom_secret"
	os.Setenv("JWT_SECRET", secret)
	// We want to verify that GenerateToken picks this up.
	// NOTE: In the current implementation (before fix), this will NOT pick it up,
	// so the token will be signed with the default key.
	// ValidateToken (before fix) will also use the default key.
	// So if we run this test BEFORE the fix:
	// 1. GenerateToken uses default key.
	// 2. ValidateToken uses default key.
	// 3. Validation succeeds.
	//
	// To verify the fix, we need to show that:
	// With the fix, GenerateToken uses "my_custom_secret".
	// If we then unset the env var (reverting to default key), ValidateToken should FAIL.

	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Verify the token is valid while the secret is still set
	_, err = ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed with custom secret set: %v", err)
	}

	// Now Unset the secret, reverting to default
	os.Unsetenv("JWT_SECRET")

	// Try to validate the token (which was signed with "my_custom_secret")
	// using the default key. It SHOULD fail if the code honored the env var.
	// If the code IGNORED the env var (current state), it signed with default,
	// and is validating with default, so it will SUCCEED (which is a failure of this test).
	_, err = ValidateToken(token)
	if err == nil {
		t.Error("Security verification failed: Token signed with custom secret should not be valid when secret is unset (default key used)")
	}
}

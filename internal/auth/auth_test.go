package auth

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt"
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

func TestJwtSecretEnv(t *testing.T) {
	// Backup original env
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Set a custom secret
	customSecret := "my_custom_secret_key_123"
	os.Setenv("JWT_SECRET", customSecret)

	// Generate a token. If the implementation respects JWT_SECRET,
	// it should use customSecret.
	tokenStr, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Verify the token was signed with the custom secret
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(customSecret), nil
	})

	// If the implementation does NOT respect JWT_SECRET (current state),
	// this manual verification with customSecret will fail.
	// But we want it to pass eventually.
	// So for now, if I run this test against the current code, it will fail.

	// However, I want to assert that the current code DOES use the env var (which is my goal).
	if err != nil || !token.Valid {
		t.Logf("Validation with custom secret failed (expected before fix): %v", err)
		// This test expects the fix to be applied.
		// So I should expect success here.
		t.Fail()
	}

	// Verify ValidateToken also works with the env var set
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("ValidateToken failed with correct env var: %v", err)
	}
	if claims.Email != "test@example.com" {
		t.Errorf("Claims mismatch")
	}

	// Now verify that if we change the env var (or unset it), validation fails
	os.Unsetenv("JWT_SECRET")
	_, err = ValidateToken(tokenStr)
	if err == nil {
		t.Error("Expected validation failure when JWT_SECRET is unset (mismatched keys), but got success")
	}
}

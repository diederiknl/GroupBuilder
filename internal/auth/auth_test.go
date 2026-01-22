package auth

import (
	"testing"
	"time"
)

func TestGenerateAndValidateToken(t *testing.T) {
	email := "test@example.com"
	role := "student"

	tokenString, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Expected no error when generating token, got %v", err)
	}

	if tokenString == "" {
		t.Fatal("Expected token string to be not empty")
	}

	claims, err := ValidateToken(tokenString)
	if err != nil {
		t.Fatalf("Expected no error when validating token, got %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestValidateToken_Invalid(t *testing.T) {
	_, err := ValidateToken("invalid-token")
	if err == nil {
		t.Error("Expected error when validating invalid token")
	}
}

func TestTokenExpiration(t *testing.T) {
	// This test is tricky because we can't easily mock time.Now() in the current implementation
	// without refactoring. However, we can inspect the claims expiry.

	email := "test@example.com"
	role := "student"

	tokenString, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(tokenString)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	// Check if expiration is roughly 24 hours from now
	expectedDuration := 24 * time.Hour
	// Allow for some execution time drift (e.g. 5 seconds)
	diff := time.Until(time.Unix(claims.ExpiresAt, 0))

	if diff > expectedDuration || diff < expectedDuration - 5*time.Second {
		t.Errorf("Expected expiration in ~24h, but got %v", diff)
	}
}

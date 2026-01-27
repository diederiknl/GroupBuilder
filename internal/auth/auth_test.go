package auth

import (
	"testing"
	"time"
)

func TestGenerateAndValidateToken(t *testing.T) {
	email := "test@example.com"
	role := "teacher"

	// Test Token Generation
	tokenString, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	if tokenString == "" {
		t.Fatal("Generated token is empty")
	}

	// Test Token Validation
	claims, err := ValidateToken(tokenString)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}

	// Check Expiration (roughly)
	if claims.ExpiresAt <= time.Now().Unix() {
		t.Error("Token is already expired")
	}
}

func TestValidateInvalidToken(t *testing.T) {
	_, err := ValidateToken("invalid.token.string")
	if err == nil {
		t.Error("Expected error for invalid token, got nil")
	}
}

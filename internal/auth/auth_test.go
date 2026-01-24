package auth

import (
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	email := "test@example.com"
	role := "student"

	tokenString, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	claims, err := ValidateToken(tokenString)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestInvalidToken(t *testing.T) {
	_, err := ValidateToken("invalid.token.string")
	if err == nil {
		t.Error("Expected error for invalid token, got nil")
	}
}

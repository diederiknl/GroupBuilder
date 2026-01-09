package auth_test

import (
	"GroupBuilder/internal/auth"
	"testing"
)

func TestTokenGenerationAndValidation(t *testing.T) {
	email := "test@example.com"
	role := "student"

	token, err := auth.GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := auth.ValidateToken(token)
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

func TestGenerateLoginLink(t *testing.T) {
    email := "test@example.com"
    link, err := auth.GenerateLoginLink(email)
    if err != nil {
        t.Fatalf("Failed to generate login link: %v", err)
    }
    if link == "" {
        t.Error("Generated login link is empty")
    }
}

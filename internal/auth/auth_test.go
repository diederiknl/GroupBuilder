package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestGenerateAndValidateToken(t *testing.T) {
	email := "test@example.com"
	role := "student"

	// Test Token Generation
	tokenString, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if tokenString == "" {
		t.Fatal("Token string is empty")
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
}

func TestValidateInvalidToken(t *testing.T) {
	// Create a token signed with a different key
	claims := &Claims{
		Email: "hack@example.com",
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign with a different key
	wrongKey := []byte("wrong_key")
	tokenString, _ := token.SignedString(wrongKey)

	_, err := ValidateToken(tokenString)
	if err == nil {
		t.Error("Expected error when validating token signed with wrong key, got nil")
	}
}

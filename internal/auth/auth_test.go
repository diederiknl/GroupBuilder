package auth

import (
	"os"
	"testing"
	"time"

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

func TestExpiredToken(t *testing.T) {
	// Create an expired token.
	// Since jwtKey is private within the package, we can access it here to sign the token manually.
	expirationTime := time.Now().Add(-1 * time.Hour)
	claims := &Claims{
		Email: "expired@example.com",
		Role:  "student",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(jwtKey)

	_, err := ValidateToken(tokenStr)
	if err == nil {
		t.Error("Expected error for expired token, got nil")
	}
}

func TestDefaultKeyFallback(t *testing.T) {
	// This test verifies the behavior relative to the JWT_SECRET environment variable.
	// 1. If JWT_SECRET is unset, the system falls back to "your_secret_key".
	// 2. If JWT_SECRET is set, the system uses the new key, so validation with the default key should fail.

	defaultKey := []byte("your_secret_key")

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: "check@example.com",
		Role:  "student",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(defaultKey)

	_, err := ValidateToken(tokenStr)

	isEnvSet := os.Getenv("JWT_SECRET") != ""

	if isEnvSet {
		// If custom key is active, validation with default key must fail.
		if err == nil {
			t.Error("Expected validation to fail when JWT_SECRET is set (key mismatch), but it passed")
		}
	} else {
		// If default key is active, validation with default key must pass.
		if err != nil {
			t.Errorf("Expected validation to pass with default key, but got error: %v", err)
		}
	}
}

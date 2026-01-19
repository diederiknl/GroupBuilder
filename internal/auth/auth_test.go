package auth

import (
	"testing"

	"github.com/golang-jwt/jwt"
)

func TestDefaultKeyIsUsedWhenEnvNotSet(t *testing.T) {
	// Manually create a token signed with "your_secret_key"
	token := jwt.New(jwt.SigningMethodHS256)
	signedString, _ := token.SignedString([]byte("your_secret_key"))

	// Validate it using the library function which should be using the default key
	_, err := ValidateToken(signedString)
	if err != nil {
		t.Errorf("Expected default key to valid manually signed token, but got error: %v", err)
	}
}

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

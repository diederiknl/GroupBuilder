package auth

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestGenerateAndValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

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

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	_, err := ValidateToken("invalid.token.string")
	if err == nil {
		t.Error("Expected error for invalid token, got nil")
	}
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	// Create an expired token manually since GenerateToken creates a valid one
	expirationTime := time.Now().Add(-1 * time.Hour)
	claims := &Claims{
		Email: "expired@example.com",
		Role:  "student",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key, err := getJwtKey()
	if err != nil {
		t.Fatalf("Failed to get jwt key: %v", err)
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	_, err = ValidateToken(tokenString)
	if err == nil {
		t.Error("Expected error for expired token, got nil")
	}
}

func TestMissingSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}

	// ValidateToken should also fail if secret is missing
	_, err = ValidateToken("some.token.here")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

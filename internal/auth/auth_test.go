package auth

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestGenerateToken_UsesEnvVar(t *testing.T) {
	expectedSecret := "test_secret_env"
	os.Setenv("JWT_SECRET", expectedSecret)
	defer os.Unsetenv("JWT_SECRET")

	tokenStr, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Verify the token was signed with the expected secret
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(expectedSecret), nil
	})

	if err != nil {
		t.Errorf("Failed to parse token with expected secret: %v", err)
	}

	if !token.Valid {
		t.Error("Token is invalid when parsed with expected secret")
	}
}

func TestValidateToken_UsesEnvVar(t *testing.T) {
	secret := "another_secret_env"
	os.Setenv("JWT_SECRET", secret)
	defer os.Unsetenv("JWT_SECRET")

	// Create a token signed with the secret
	claims := &Claims{
		Email: "user@example.com",
		Role:  "teacher",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Failed to sign token manually: %v", err)
	}

	// Validate using the auth package function
	parsedClaims, err := ValidateToken(tokenStr)
	if err != nil {
		t.Errorf("ValidateToken failed with env var secret: %v", err)
	}
	if parsedClaims == nil || parsedClaims.Email != "user@example.com" {
		t.Error("Parsed claims do not match expected")
	}
}

func TestFallbackBehavior(t *testing.T) {
	// Ensure JWT_SECRET is unset
	os.Unsetenv("JWT_SECRET")

	// The default fallback key from auth.go
	expectedKey := []byte("your_secret_key")

	tokenStr, err := GenerateToken("fallback@example.com", "student")
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Verify the token was signed with the fallback key
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return expectedKey, nil
	})

	if err != nil {
		t.Errorf("Failed to parse token with fallback key: %v", err)
	}
	if !token.Valid {
		t.Error("Token signed with fallback key is invalid")
	}
}

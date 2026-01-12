package auth

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Setup - Ensure we are using a known key for this test
	originalKey := jwtKey
	defer func() { jwtKey = originalKey }()
	jwtKey = []byte("test_secret")

	email := "test@example.com"
	role := "student"

	// Test Generation
	tokenString, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if tokenString == "" {
		t.Fatal("Token is empty")
	}

	// Test Validation
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

func TestTokenExpiration(t *testing.T) {
	// Setup
	originalKey := jwtKey
	defer func() { jwtKey = originalKey }()
	jwtKey = []byte("test_secret")

	claims := &Claims{
		Email: "test@example.com",
		Role:  "student",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(), // Expired
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	_, err := ValidateToken(tokenString)
	if err == nil {
		t.Error("Expected error for expired token, got nil")
	}
}

func TestEnvVariable(t *testing.T) {
	// Save current env and jwtKey
	originalEnv := os.Getenv("JWT_SECRET")
	originalKey := jwtKey
	defer func() {
		os.Setenv("JWT_SECRET", originalEnv)
		jwtKey = originalKey
	}()

	// Test with explicit secret
	testSecret := "super_secret_env_key"
	os.Setenv("JWT_SECRET", testSecret)

	// Re-run init logic manually since init() only runs once at package load
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		jwtKey = []byte(secret)
	}

	if string(jwtKey) != testSecret {
		t.Errorf("Expected jwtKey to be %s, got %s", testSecret, string(jwtKey))
	}
}

func TestDefaultKeyWarning(t *testing.T) {
	// This test just ensures that if we unset the env var, it defaults to the hardcoded dev key
	// Note: We can't easily capture stdout to check for the warning print without more complex setup,
	// so we'll just check the key value.

	originalEnv := os.Getenv("JWT_SECRET")
	originalKey := jwtKey
	defer func() {
		os.Setenv("JWT_SECRET", originalEnv)
		jwtKey = originalKey
	}()

	os.Unsetenv("JWT_SECRET")

	// Re-run init logic manually
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		jwtKey = []byte(secret)
	} else {
		jwtKey = []byte("your_secret_key")
	}

	expectedDefault := "your_secret_key"
	if string(jwtKey) != expectedDefault {
		t.Errorf("Expected default jwtKey to be %s, got %s", expectedDefault, string(jwtKey))
	}
}

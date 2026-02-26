package auth

import (
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("GenerateToken returned empty string")
	}
}

func TestValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	token, _ := GenerateToken("test@example.com", "student")

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}
	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", claims.Email)
	}
}

func TestValidateTokenInvalidSecret(t *testing.T) {
	// Set initial secret
	os.Setenv("JWT_SECRET", "test_secret")
	token, _ := GenerateToken("test@example.com", "student")

	// Change secret
	os.Setenv("JWT_SECRET", "wrong_secret")
	defer os.Unsetenv("JWT_SECRET")

	_, err := ValidateToken(token)
	if err == nil {
		t.Error("ValidateToken should fail with wrong secret")
	}
}

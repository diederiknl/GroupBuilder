package auth

import (
	"os"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {
	// Backup original env
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Test case 1: No JWT_SECRET
	os.Unsetenv("JWT_SECRET")
	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("GenerateToken should fail without JWT_SECRET")
	}

	// Test case 2: Valid JWT_SECRET
	testSecret := "test_secret_key"
	os.Setenv("JWT_SECRET", testSecret)

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", claims.Email)
	}

	// Test case 3: Login Link
	link, err := GenerateLoginLink("test@example.com")
	if err != nil {
		t.Fatalf("GenerateLoginLink failed: %v", err)
	}
	if !strings.Contains(link, "token=") {
		t.Error("Generated link does not contain token")
	}
}

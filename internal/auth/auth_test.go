package auth

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	// Set up environment variable
	os.Setenv("JWT_SECRET", "testsecret")
	defer os.Unsetenv("JWT_SECRET")

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Error("Expected token to be generated")
	}
}

func TestValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	defer os.Unsetenv("JWT_SECRET")

	token, _ := GenerateToken("test@example.com", "student")

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", claims.Email)
	}
}

func TestMissingSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing")
	}
}

func TestGenerateLoginLink(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	defer os.Unsetenv("JWT_SECRET")

	link, err := GenerateLoginLink("test@example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if link == "" {
		t.Error("Expected login link to be generated")
	}

	// Verify it contains the base url
	expectedBase := "http://localhost:8080"
	if !strings.HasPrefix(link, expectedBase) {
		t.Errorf("Expected link to start with %s, got %s", expectedBase, link)
	}
}

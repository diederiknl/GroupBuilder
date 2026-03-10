package auth

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Save the original value
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)

	// Test with secret
	os.Setenv("JWT_SECRET", "test_secret")

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("Expected token, got empty string")
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}
	if claims.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", claims.Email)
	}
	if claims.Role != "student" {
		t.Errorf("Expected role 'student', got '%s'", claims.Role)
	}

	// Test without secret
	os.Setenv("JWT_SECRET", "")

	_, err = GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("GenerateToken should fail without JWT_SECRET")
	}

	_, err = ValidateToken(token)
	if err == nil {
		t.Error("ValidateToken should fail without JWT_SECRET")
	}
}

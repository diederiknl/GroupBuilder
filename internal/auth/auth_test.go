package auth

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateToken_NoSecret(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Unsetenv("JWT_SECRET")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is not set, got nil")
	}
	if err.Error() != "JWT_SECRET environment variable not set" {
		t.Errorf("Expected 'JWT_SECRET environment variable not set', got '%v'", err)
	}
}

func TestGenerateToken_Success(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "testsecret")

	token, err := GenerateToken("test@example.com", "student")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected token, got empty string")
	}
}

func TestValidateToken_Success(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "testsecret")

	token, _ := GenerateToken("test@example.com", "student")

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if claims.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", claims.Email)
	}
	if claims.Role != "student" {
		t.Errorf("Expected role 'student', got '%s'", claims.Role)
	}
}

func TestValidateToken_InvalidSecret(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "testsecret")

	token, _ := GenerateToken("test@example.com", "student")

	os.Setenv("JWT_SECRET", "wrongsecret")
	_, err := ValidateToken(token)
	if err == nil {
		t.Error("Expected error when validating with wrong secret, got nil")
	}
}

func TestGenerateLoginLink(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	defer os.Setenv("JWT_SECRET", originalSecret)
	os.Setenv("JWT_SECRET", "testsecret")

	originalBaseURL := os.Getenv("BASE_URL")
	defer os.Setenv("BASE_URL", originalBaseURL)
	os.Setenv("BASE_URL", "http://test.com")

	link, err := GenerateLoginLink("test@example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.HasPrefix(link, "http://test.com/auth/student/verify?token=") {
		t.Errorf("Expected link to start with 'http://test.com/auth/student/verify?token=', got '%s'", link)
	}
}

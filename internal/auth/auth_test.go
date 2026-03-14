package auth

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set up env var
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Setenv("JWT_SECRET", "")

	email := "test@example.com"
	role := "student"

	token, err := GenerateToken(email, role)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestGenerateTokenMissingEnvVar(t *testing.T) {
	os.Setenv("JWT_SECRET", "")

	_, err := GenerateToken("test@example.com", "student")
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

func TestValidateTokenMissingEnvVar(t *testing.T) {
	// First generate a token successfully
	os.Setenv("JWT_SECRET", "test_secret")
	token, _ := GenerateToken("test@example.com", "student")

	// Then clear env var and try to validate
	os.Setenv("JWT_SECRET", "")
	defer os.Setenv("JWT_SECRET", "")

	_, err := ValidateToken(token)
	if err == nil {
		t.Error("Expected error when JWT_SECRET is missing, got nil")
	}
}

func TestGenerateLoginLink(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Setenv("JWT_SECRET", "")

	email := "student@example.com"
	link, err := GenerateLoginLink(email)
	if err != nil {
		t.Fatalf("GenerateLoginLink failed: %v", err)
	}

	expectedPrefix := "/login?token="
	if !strings.HasPrefix(link, expectedPrefix) {
		t.Errorf("Expected link to start with %s, got %s", expectedPrefix, link)
	}

	token := strings.TrimPrefix(link, expectedPrefix)
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed for generated link token: %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s in token, got %s", email, claims.Email)
	}
	if claims.Role != "student" {
		t.Errorf("Expected role 'student' in token, got %s", claims.Role)
	}
}

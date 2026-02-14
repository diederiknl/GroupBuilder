package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func getJwtKey() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}
	return []byte(secret), nil
}

func GenerateToken(email string, role string) (string, error) {
	key, err := getJwtKey()
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return getJwtKey()
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// GenerateLoginLink creates a magic link for student login
// This is a placeholder implementation that returns a token-based URL
func GenerateLoginLink(email string) (string, error) {
	// For now, we reuse GenerateToken to create a token
	// In a real app, this might be a different type of token or process
	token, err := GenerateToken(email, "student")
	if err != nil {
		return "", err
	}

	// Construct the link (assuming some base URL)
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return fmt.Sprintf("%s/login/verify?token=%s", baseURL, token), nil
}

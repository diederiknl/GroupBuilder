package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// getJwtKey retrieves the JWT secret from the environment variable.
// If not set, it defaults to the legacy hardcoded key (for backward compatibility).
func getJwtKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return []byte("your_secret_key")
	}
	return []byte(secret)
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(email string, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJwtKey())
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return getJwtKey(), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// GenerateLoginLink creates a magic link for student login
func GenerateLoginLink(email string) (string, error) {
	token, err := GenerateToken(email, "student")
	if err != nil {
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	return fmt.Sprintf("%s/login/verify?token=%s", baseURL, token), nil
}

package auth

import (
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
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable not set")
	}
	return []byte(key), nil
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

// GenerateLoginLink generates a token for student login
func GenerateLoginLink(email string) (string, error) {
	// For now, we reuse GenerateToken with "student" role for login link
	// In a real app, this might be a short-lived token just for login verification
	return GenerateToken(email, "student")
}

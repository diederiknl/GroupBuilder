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
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}
	return []byte(key), nil
}

func GenerateToken(email string, role string) (string, error) {
	jwtKey, err := getJwtKey()
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
	return token.SignedString(jwtKey)
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

func GenerateLoginLink(email string) (string, error) {
	token, err := GenerateToken(email, "student")
	if err != nil {
		return "", err
	}
	// In a real app, this should be a configurable base URL
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return fmt.Sprintf("%s/login?token=%s", baseURL, token), nil
}

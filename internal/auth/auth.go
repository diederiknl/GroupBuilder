package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// jwtKey is removed in favor of getJwtKey()

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func getJwtKey() ([]byte, error) {
	// Security: Retrieve the secret key from environment variables to prevent hardcoding secrets in the codebase.
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}
	return []byte(secret), nil
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
	token, err := GenerateToken(email, "login")
	if err != nil {
		return "", err
	}
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return baseURL + "/login?token=" + token, nil
}

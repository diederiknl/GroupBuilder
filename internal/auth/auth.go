package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// getJwtKey retrieves the JWT secret from the environment variable.
func getJwtKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// This serves as a runtime safeguard.
		panic("JWT_SECRET environment variable not set")
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

// GenerateLoginLink generates a login link for the given email.
// This is currently a stub implementation.
func GenerateLoginLink(email string) (string, error) {
	// TODO: Implement actual login link generation logic
	// For now, we'll return a dummy link.
	return fmt.Sprintf("http://localhost:8080/login?email=%s&token=dummy_token", email), nil
}

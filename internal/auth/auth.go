package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func getJwtKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
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

func GenerateLoginLink(email string) (string, error) {
	// Stub implementation
	// In a real implementation, this would generate a secure token and link
	// For now, we return a dummy link or error as it's not fully implemented
	return "http://localhost:8080/auth/login?token=dummy", nil
}

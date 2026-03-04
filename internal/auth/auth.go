package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey []byte

// InitKey initializes the JWT secret key from the environment variable.
func InitKey() error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return errors.New("JWT_SECRET environment variable is required but not set")
	}
	jwtKey = []byte(secret)
	return nil
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
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

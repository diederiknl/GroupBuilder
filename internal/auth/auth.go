package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func getJWTKey() ([]byte, error) {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return nil, errors.New("JWT_SECRET environment variable is not set")
	}
	return []byte(key), nil
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(email string, role string) (string, error) {
	key, err := getJWTKey()
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
	key, err := getJWTKey()
	if err != nil {
		return nil, err
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
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
	return fmt.Sprintf("/login?token=%s", token), nil
}

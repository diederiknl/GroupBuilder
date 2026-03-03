package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey []byte

func InitKey() {
	secret := os.Getenv("JWT_SECRET")
	if secret != "" {
		jwtKey = []byte(secret)
	}
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateLoginLink(email string) (string, error) {
	return "http://localhost:8080/auth/student/verify?token=stub_token_for_" + email, nil
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

package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("WARNING: JWT_SECRET environment variable is not set. Using default development key.")
		jwtKey = []byte("dev-secret-key")
	} else {
		jwtKey = []byte(secret)
	}
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

func GenerateLoginLink(email string) (string, error) {
	// For now, this just generates a token.
	// In a real app, this might generate a unique link with a temporary token
	token, err := GenerateToken(email, "student_login_pending")
	if err != nil {
		return "", err
	}
	// Assuming the link structure, this might need to be adjusted based on frontend/requirements
	return "http://localhost:8080/login?token=" + token, nil
}

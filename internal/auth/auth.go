package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func getJWTKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return []byte("your_secret_key")
	}
	return []byte(key)
}

func init() {
	if os.Getenv("JWT_SECRET") == "" {
		log.Println("WARNING: JWT_SECRET environment variable is not set. Using default insecure key.")
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
	return token.SignedString(getJWTKey())
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTKey(), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// GenerateLoginLink generates a secure random token for a login link
func GenerateLoginLink(email string) (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)

	// In a real app, you'd constructing a full URL here, e.g.:
	// baseURL := os.Getenv("BASE_URL")
	// return fmt.Sprintf("%s/auth/verify?token=%s&email=%s", baseURL, token, email), nil

	// For now, returning just the token string or a dummy link format
	return fmt.Sprintf("http://localhost:8080/auth/verify?token=%s", token), nil
}

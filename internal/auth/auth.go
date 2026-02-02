package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Default to a known key for development if not set, but warn about it.
// In production, JWT_SECRET must be set.
var defaultJWTKey = []byte("your_secret_key")

func init() {
	if os.Getenv("JWT_SECRET") == "" {
		log.Println("WARNING: JWT_SECRET environment variable is not set. Using default insecure key for development.")
	}
}

func getJWTKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return defaultJWTKey
	}
	return []byte(key)
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

func GenerateLoginLink(email string) (string, error) {
	token, err := GenerateToken(email, "student")
	if err != nil {
		return "", err
	}
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return fmt.Sprintf("%s/auth/login?token=%s", baseURL, token), nil
}

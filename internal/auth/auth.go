package auth

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func getJwtKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback to legacy key if env var is not set
		return []byte("your_secret_key")
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

func GenerateLoginLink(email string) (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)

	// In a real app, this would construct a full URL like:
	// return fmt.Sprintf("%s/login?token=%s", baseURL, token), nil
	// For now, returning just the token string to satisfy the interface,
	// or a mock link as the handler expects.
	// Looking at handler usage: link, err := auth.GenerateLoginLink(req.Email)
	// The handler just logs it or sends it.

	// Constructing a basic link format
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return baseURL + "/auth/verify?token=" + token + "&email=" + email, nil
}

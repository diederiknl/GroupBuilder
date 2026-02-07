package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func getJwtKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return []byte("your_secret_key")
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

// GenerateLoginLink generates a link (or just a token for now) for student login.
// In a real app, this would generate a unique token, store it, and return a link.
// For this fix, we'll return a dummy link.
func GenerateLoginLink(email string) (string, error) {
	// Reusing GenerateToken for simplicity in this hotfix context,
	// though typically this would be a one-time use token.
	token, err := GenerateToken(email, "student")
	if err != nil {
		return "", err
	}
	// Assuming a frontend route or API endpoint verification
	return "http://localhost:8080/auth/student/verify?token=" + token, nil
}

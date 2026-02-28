package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey []byte

func InitKey() {
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

func getJwtKey() []byte {
	if len(jwtKey) == 0 {
		InitKey()
	}
	return jwtKey
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

// GenerateLoginLink generates a relative login link for a student
func GenerateLoginLink(email string) (string, error) {
	// In a real application, you would create a secure, randomly generated token,
	// save it to the database with an expiration time, and construct a link.
	// For now, we'll return a stub link.
	return "/auth/student/login?token=stub_token_for_" + email, nil
}

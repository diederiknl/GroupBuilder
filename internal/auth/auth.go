package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey []byte

func init() {
	loadKey()
}

func loadKey() {
	key, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		log.Println("WARNING: JWT_SECRET environment variable is not set. Using default insecure key. DO NOT USE IN PRODUCTION.")
		jwtKey = []byte("default_insecure_secret_for_dev_only")
	} else {
		jwtKey = []byte(key)
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

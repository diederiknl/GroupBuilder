package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// getJWTKey retrieves the JWT secret key from the environment variable.
// If the variable is not set, it defaults to a hardcoded key for development purposes,
// and logs a warning.
func getJWTKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		log.Println("WARNING: JWT_SECRET environment variable not set. Using insecure default key for development.")
		return []byte("your_secret_key")
	}
	return []byte(key)
}

// getBaseURL retrieves the base URL for the application from the environment.
// Defaults to http://localhost:8080 if not set.
func getBaseURL() string {
	url := os.Getenv("BASE_URL")
	if url == "" {
		return "http://localhost:8080"
	}
	return url
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

// GenerateLoginLink generates a magic link token for student login.
func GenerateLoginLink(email string) (string, error) {
	// Create a token that expires in 15 minutes
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: email,
		Role:  "student_login",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTKey())
	if err != nil {
		return "", err
	}

	baseURL := getBaseURL()
	return fmt.Sprintf("%s/auth/student/verify?token=%s", baseURL, tokenString), nil
}

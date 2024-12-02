package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginLink struct {
	Token     string
	ExpiresAt time.Time
}

func GenerateLoginLink(email string) (LoginLink, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return LoginLink{}, err
	}

	return LoginLink{
		Token:     base64.URLEncoding.EncodeToString(token),
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}, nil
}

func ValidateLoginLink(token string) bool {
	// Implementeer de logica om de token te valideren en te controleren of deze nog geldig is
	// Dit zal waarschijnlijk een database-lookup vereisen
	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package auth

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	signingKey = "b4e81c3a0d09874e5d71bf2a55b3c1e70f3d04a44c7c414bda872ef8f22a7b7f"
	tokenTTL   = 12 * time.Hour
	salt       = "7f2cbe9876a1a3cfe9952ebccda7e144"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userID,
	})

	return token.SignedString([]byte(signingKey))
}

func GenerateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

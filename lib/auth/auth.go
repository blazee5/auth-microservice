package auth

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	signingKey        = "b4e81c3a0d09874e5d71bf2a55b3c1e70f3d04a44c7c414bda872ef8f22a7b7f"
	refreshSigningKey = "YxrN2XqOeEmruOT0XLLNZwPYe9bjXaUx"
	accessTokenTTL    = 12 * time.Hour
	refreshTokenTTL   = 24 * time.Hour * 30
	salt              = "7f2cbe9876a1a3cfe9952ebccda7e144"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

type Tokens struct {
	Access  string
	Refresh string
}

func GenerateTokens(userID string) (Tokens, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userID,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userID,
	})

	access, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return Tokens{}, err
	}

	refresh, err := refreshToken.SignedString([]byte(refreshSigningKey))

	if err != nil {
		return Tokens{}, err
	}

	tokens := Tokens{
		Access:  access,
		Refresh: refresh,
	}

	return tokens, nil
}

func GenerateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func ParseToken(token string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshSigningKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok {
		return "", err
	}

	return claims.UserID, nil
}

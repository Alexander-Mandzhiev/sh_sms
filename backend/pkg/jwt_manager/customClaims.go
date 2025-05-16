package jwt_manager

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type CustomClaims struct {
	UserID      uuid.UUID `json:"user_id"`
	ClientID    uuid.UUID `json:"client_id"`
	AppID       int       `json:"app_id"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
	TokenType   TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

func HashToken(token string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(token)))
}

func Generate(claims CustomClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

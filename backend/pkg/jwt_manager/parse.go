package jwt_manager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func Parse(tokenString, secret string, expectedType TokenType) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("token parse failed: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	if claims.TokenType != expectedType {
		return nil, fmt.Errorf("invalid token type: expected %s", expectedType)
	}

	return claims, nil
}

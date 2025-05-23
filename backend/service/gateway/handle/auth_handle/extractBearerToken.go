package auth_handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) extractTokenFromCookie(c *gin.Context, tokenType string) (string, error) {
	cookieName := ""
	switch tokenType {
	case "access":
		cookieName = "access_token"
	case "refresh":
		cookieName = "refresh_token"
	default:
		return "", fmt.Errorf("invalid token type")
	}

	token, err := c.Cookie(cookieName)
	if err != nil {
		return "", fmt.Errorf("cookie %s not found: %w", cookieName, err)
	}

	if token == "" {
		return "", fmt.Errorf("empty %s token", cookieName)
	}

	return token, nil
}

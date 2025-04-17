package repository

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateSecureSecret() (string, error) {
	const length = 32
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate secret: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

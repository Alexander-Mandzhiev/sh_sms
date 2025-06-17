package utils

import "strings"

func NormalizeEmail(email string) string {
	return strings.ToLower(TrimSpace(email))
}

package utils

import "strings"

func IsValidPhone(phone string) bool {
	if phone == "" {
		return false
	}
	if !strings.HasPrefix(phone, "+") {
		return false
	}
	return len(phone) >= 8 && len(phone) <= 20
}

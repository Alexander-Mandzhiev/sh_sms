package utils

import "strings"

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func TrimSpacePointer(s *string) *string {
	if s == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*s)
	return &trimmed
}

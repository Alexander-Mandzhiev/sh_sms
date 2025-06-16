package private_school_models

import "strings"

func cleanPhone(phone string) string {
	if phone == "" {
		return ""
	}

	var b strings.Builder
	if strings.HasPrefix(phone, "+") {
		b.WriteByte('+')
		phone = phone[1:]
	}

	for _, r := range phone {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}

	return b.String()
}

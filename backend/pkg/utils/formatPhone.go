package utils

func FormatPhone(phone string) string {
	cleaned := CleanPhone(phone)
	if cleaned == "" {
		return ""
	}

	if len(cleaned) < 7 {
		return ""
	}
	return "+" + cleaned
}

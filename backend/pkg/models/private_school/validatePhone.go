package private_school_models

func isValidPhone(phone string) bool {
	if phone == "" {
		return true
	}

	if len(phone) < 2 || len(phone) > 16 {
		return false
	}

	if phone[0] != '+' {
		return false
	}

	for _, c := range phone[1:] {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

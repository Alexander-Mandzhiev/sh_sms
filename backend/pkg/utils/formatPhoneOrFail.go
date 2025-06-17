package utils

import "errors"

func FormatPhoneOrFail(phone string) (string, error) {
	cleaned := CleanPhone(phone)
	if cleaned == "" {
		return "", errors.New("empty phone number")
	}

	if len(cleaned) < 8 {
		return "", errors.New("phone number too short")
	}

	if len(cleaned) > 20 {
		return "", errors.New("phone number too long")
	}

	return "+" + cleaned, nil
}

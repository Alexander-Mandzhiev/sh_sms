package private_school_models

import (
	"regexp"
	"strings"
)

func isValidEmail(email string) bool {
	if email == "" {
		return true
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	local, domain := parts[0], parts[1]
	if local == "" || domain == "" {
		return false
	}

	if !strings.Contains(domain, ".") {
		return false
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

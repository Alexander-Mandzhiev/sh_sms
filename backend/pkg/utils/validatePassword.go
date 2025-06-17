package utils

import (
	"backend/service/constants"
	"fmt"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("%w: password must be at least 8 characters", constants.ErrInvalidArgument)
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return fmt.Errorf("%w: password must contain uppercase letters", constants.ErrInvalidArgument)
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return fmt.Errorf("%w: password must contain digits", constants.ErrInvalidArgument)
	}

	return nil
}

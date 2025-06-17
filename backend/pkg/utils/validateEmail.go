package utils

import (
	"backend/service/constants"
	"fmt"
)

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("%w: email is required", constants.ErrInvalidArgument)
	}
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("%w: invalid email format", constants.ErrInvalidArgument)
	}
	return nil
}

package utils

import (
	"backend/service/constants"
	"fmt"
)

func ValidateString(code string, length int) error {
	if code == "" {
		return fmt.Errorf("%w: code required", constants.ErrInvalidArgument)
	}
	if len(code) > length {
		return fmt.Errorf("%w: code too long", constants.ErrInvalidArgument)
	}
	return nil
}

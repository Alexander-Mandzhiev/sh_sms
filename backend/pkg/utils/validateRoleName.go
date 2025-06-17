package utils

import (
	"backend/service/constants"
	"fmt"
)

func ValidateRoleName(name string, length int) error {
	if name == "" || len(name) > length {
		return fmt.Errorf("%w: name is required", constants.ErrInvalidArgument)
	}
	return nil
}

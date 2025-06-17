package utils

import (
	"backend/service/constants"
	"fmt"
)

func ValidateRoleLevel(level int) error {
	if level < 0 {
		return fmt.Errorf("%w: invalid level", constants.ErrInvalidArgument)
	}
	return nil
}

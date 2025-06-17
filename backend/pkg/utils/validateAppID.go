package utils

import (
	"backend/service/constants"
	"fmt"
)

func ValidateAppID(appID int) error {
	if appID <= 0 {
		return fmt.Errorf("%w: appID must be positive", constants.ErrInvalidAppId)
	}

	return nil
}

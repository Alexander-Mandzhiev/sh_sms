package service

import (
	"backend/service/constants"
	"fmt"
)

func validateID(id int) error {
	if id <= 0 {
		return fmt.Errorf("%w: %d", constants.ErrInvalidID, id)
	}
	return nil
}

func validateCode(code string, maxLength int) error {
	if code == "" {
		return constants.ErrEmptyCode
	}
	if len(code) > maxLength {
		return fmt.Errorf("%w: code exceeds max length %d", constants.ErrInvalidCode, maxLength)
	}
	return nil
}

func validateName(name string, maxLength int) error {
	if name == "" {
		return constants.ErrEmptyName
	}
	if len(name) > maxLength {
		return fmt.Errorf("%w: name length %d > %d",
			constants.ErrInvalidName, len(name), maxLength)
	}
	return nil
}

package service

import (
	"backend/service/apps/app_manager/handle"
	"fmt"
)

func validateID(id int) error {
	if id <= 0 {
		return fmt.Errorf("%w: %d", handle.ErrInvalidID, id)
	}
	return nil
}

func validateCode(code string, maxLength int) error {
	if code == "" {
		return handle.ErrEmptyCode
	}
	if len(code) > maxLength {
		return fmt.Errorf("%w: code exceeds max length %d", handle.ErrInvalidCode, maxLength)
	}
	return nil
}

func validateName(name string, maxLength int) error {
	if name == "" {
		return handle.ErrEmptyName
	}
	if len(name) > maxLength {
		return fmt.Errorf("%w: name length %d > %d",
			handle.ErrInvalidName, len(name), maxLength)
	}
	return nil
}

func validatePagination(page, count int) error {
	if page < 1 || count < 1 {
		return fmt.Errorf("%w: page=%d count=%d",
			handle.ErrInvalidPagination, page, count)
	}

	return nil
}

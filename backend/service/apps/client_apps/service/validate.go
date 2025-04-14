package service

import (
	"github.com/google/uuid"
)

func validateClientID(clientID string) error {
	if clientID == "" {
		return ErrInvalidArgument
	}
	if _, err := uuid.Parse(clientID); err != nil {
		return ErrInvalidArgument
	}
	return nil
}

func validateAppID(appID int32) error {
	if appID <= 0 {
		return ErrInvalidArgument
	}
	return nil
}

func validatePagination(page, count int32) error {
	if page <= 0 || count <= 0 {
		return ErrInvalidArgument
	}
	return nil
}

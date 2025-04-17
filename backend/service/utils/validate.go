package utils

import (
	"backend/service/apps/constants"
	"github.com/google/uuid"
)

func ValidateClientID(clientID string) error {
	if clientID == "" {
		return constants.ErrInvalidArgument
	}
	if _, err := uuid.Parse(clientID); err != nil {
		return constants.ErrInvalidArgument
	}
	return nil
}

func ValidateAppID(appID int) error {
	if appID <= 0 {
		return constants.ErrInvalidAppId
	}
	return nil
}

func ValidatePagination(page, count int) error {
	if page <= 0 || count <= 0 {
		return constants.ErrInvalidArgument
	}
	if count > 1000 {
		return constants.ErrInvalidArgument
	}
	return nil
}

func IsValidSecretType(secretType string) bool {
	return secretType == "access" || secretType == "refresh"
}

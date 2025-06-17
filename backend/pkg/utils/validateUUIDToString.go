package utils

import (
	"backend/service/constants"
	"github.com/google/uuid"
)

func ValidateUUIDToString(id string) error {
	parsedUUID, err := uuid.Parse(id)
	if err != nil || parsedUUID == uuid.Nil {
		return constants.ErrInvalidArgument
	}
	return nil
}

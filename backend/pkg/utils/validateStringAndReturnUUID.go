package utils

import (
	"backend/service/constants"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func ValidateStringAndReturnUUID(input string) (uuid.UUID, error) {
	if input == "" {
		return uuid.Nil, fmt.Errorf("%w: empty UUID", constants.ErrInvalidArgument)
	}

	id, err := uuid.Parse(input)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if !IsValidUUID(id) {
		return uuid.Nil, errors.New("invalid UUID version")
	}
	return id, nil
}

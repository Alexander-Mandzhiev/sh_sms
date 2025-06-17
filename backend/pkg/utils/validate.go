package utils

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateUUID(id uuid.UUID) error {
	if id == uuid.Nil {
		return constants.ErrInvalidArgument
	}
	return nil
}

func ValidatePagination(page, count int) error {
	if page <= 0 || count <= 0 {
		return constants.ErrInvalidArgument
	}
	if count > 10000 {
		return constants.ErrInvalidArgument
	}
	return nil
}

func IsValidSecretType(secretType string) bool {
	return secretType == "access" || secretType == "refresh"
}

func ValidateRotationHistory(h *models.RotationHistory) error {
	if h == nil {
		return errors.New("nil rotation history")
	}

	if _, err := ValidateStringAndReturnUUID(h.ClientID); err != nil {
		return fmt.Errorf("invalid client_id: %w", err)
	}

	if h.AppID <= 0 {
		return errors.New("invalid app_id")
	}

	if !IsValidSecretType(h.SecretType) {
		return fmt.Errorf("invalid secret_type: %s", h.SecretType)
	}

	if h.RotatedAt.IsZero() {
		return errors.New("zero rotated_at")
	}

	return nil
}

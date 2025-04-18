package utils

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"errors"
	"fmt"
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

func ValidateRotationHistory(h *models.RotationHistory) error {
	if h == nil {
		return errors.New("nil rotation history")
	}

	if err := ValidateClientID(h.ClientID); err != nil {
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

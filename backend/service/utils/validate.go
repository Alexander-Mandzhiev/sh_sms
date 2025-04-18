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
func ValidatePasswordPolicy(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("%w: password is required", constants.ErrInvalidArgument)
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("%w: email is required", constants.ErrInvalidArgument)
	}
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("%w: invalid email format", constants.ErrInvalidArgument)
	}
	return nil
}

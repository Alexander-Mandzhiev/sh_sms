package utils

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateUUID(id uuid.UUID) error {
	if id == uuid.Nil {
		return constants.ErrInvalidArgument
	}
	return nil
}

func ValidateUUIDToString(id string) error {
	parsedUUID, err := uuid.Parse(id)
	if err != nil || parsedUUID == uuid.Nil {
		return constants.ErrInvalidArgument
	}
	return nil
}

func ValidateAppID(appID int) error {
	if appID <= 0 {
		return fmt.Errorf("%w: appID must be positive", constants.ErrInvalidAppId)
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

	if _, err := ValidateAndReturnUUID(h.ClientID); err != nil {
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
		return fmt.Errorf("%w: password must be at least 8 characters", constants.ErrInvalidPassword)
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("%w: password must be at least 8 characters", constants.ErrInvalidArgument)
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return fmt.Errorf("%w: password must contain uppercase letters", constants.ErrInvalidArgument)
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return fmt.Errorf("%w: password must contain digits", constants.ErrInvalidArgument)
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

func ValidatePhone(phone string) error {
	if phone == "" {
		return nil
	}
	if len(phone) < 5 || !strings.HasPrefix(phone, "+") {
		return errors.New("phone must start with '+' and have at least 5 digits")
	}
	return nil
}

func ValidateAndReturnUUID(input string) (uuid.UUID, error) {
	if input == "" {
		return uuid.Nil, fmt.Errorf("%w: empty UUID", constants.ErrInvalidArgument)
	}

	id, err := uuid.Parse(input)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	return id, nil
}

func ValidateRoleName(name string, length int) error {
	if name == "" || len(name) > length {
		return fmt.Errorf("%w: name is required", constants.ErrInvalidArgument)
	}
	return nil
}

func ValidateRoleLevel(level int) error {
	if level < 0 {
		return fmt.Errorf("%w: invalid level", constants.ErrInvalidArgument)
	}
	return nil
}

func ValidateString(code string, length int) error {
	if code == "" {
		return fmt.Errorf("%w: code required", constants.ErrInvalidArgument)
	}
	if len(code) > length {
		return fmt.Errorf("%w: code too long", constants.ErrInvalidArgument)
	}
	return nil
}

func ValidateWebsite(website string, length int) error {
	if err := ValidateString(website, length); err != nil {
		return err
	}

	u, err := url.Parse(website)
	if err != nil {
		return fmt.Errorf("invalid URL format")
	}
	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("URL must contain scheme and host")
	}
	return nil
}

package teachers_service

import (
	private_school_models "backend/pkg/models/private_school"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) handleRepoError(logger *slog.Logger, err error, operation string) error {
	switch {
	case errors.Is(err, private_school_models.ErrDuplicateTeacher):
		logger.Warn("duplicate detected", "operation", operation, "error", err)
		return err

	case errors.Is(err, private_school_models.ErrInvalidClient):
		logger.Warn("invalid client reference", "operation", operation, "error", err)
		return err

	case errors.Is(err, private_school_models.ErrCreateFailed):
		logger.Error("operation failed unexpectedly", "operation", operation, "error", err)
		return fmt.Errorf("%s: %w", operation, err)

	default:
		logger.Error("operation failed", "operation", operation, "error", err)
		return fmt.Errorf("%s: %w", operation, err)
	}
}

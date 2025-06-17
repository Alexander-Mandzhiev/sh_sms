package teachers_service

import (
	"backend/pkg/models/teacher"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) handleRepoError(logger *slog.Logger, err error, operation string) error {
	switch {
	case errors.Is(err, teachers_models.ErrDuplicateTeacher):
		logger.Warn("duplicate detected", "operation", operation, "error", err)
		return err

	case errors.Is(err, teachers_models.ErrInvalidClient):
		logger.Warn("invalid client reference", "operation", operation, "error", err)
		return err

	case errors.Is(err, teachers_models.ErrCreateFailed):
		logger.Error("operation failed unexpectedly", "operation", operation, "error", err)
		return fmt.Errorf("%s: %w", operation, err)

	default:
		logger.Error("operation failed", "operation", operation, "error", err)
		return fmt.Errorf("%s: %w", operation, err)
	}
}

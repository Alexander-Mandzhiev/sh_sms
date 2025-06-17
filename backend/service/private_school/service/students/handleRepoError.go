package students_service

import "C"
import (
	"backend/pkg/models/students"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

func (s *Service) handleRepoError(err error, operation string, ctx ...interface{}) error {
	logCtx := []interface{}{"operation", operation, "error", err}
	logCtx = append(logCtx, ctx...)

	switch {
	case errors.Is(err, students_models.ErrDuplicateContract):
		s.logger.Warn("duplicate contract detected", logCtx...)
		return fmt.Errorf("%s: %w", operation, err)

	case errors.Is(err, students_models.ErrInvalidClientID):
		s.logger.Warn("invalid client reference", logCtx...)
		return fmt.Errorf("%s: %w", operation, err)

	case errors.Is(err, students_models.ErrCreateFailed):
		s.logger.Error("create operation failed", logCtx...)
		return fmt.Errorf("%s: %w", operation, err)

	case errors.Is(err, students_models.ErrStudentNotFound):
		s.logger.Warn("student not found", "operation", operation, "error", err)
		return err

	case errors.Is(err, students_models.ErrUpdateFailed):
		s.logger.Error("update operation failed", logCtx...)
		return fmt.Errorf("%s: %w", operation, err)

	case errors.Is(err, context.Canceled):
		s.logger.Warn("request canceled", logCtx...)
		return err

	case errors.Is(err, context.DeadlineExceeded):
		s.logger.Warn("request timed out", logCtx...)
		return status.Error(codes.DeadlineExceeded, "request timed out")

	default:
		s.logger.Error("unexpected repository error", append(logCtx, "stack", string(debug.Stack()))...)
		return fmt.Errorf("%s: %w", operation, err)
	}
}

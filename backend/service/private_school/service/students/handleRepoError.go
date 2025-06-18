package students_service

import (
	"backend/pkg/models/students"
	"context"
	"errors"
	"runtime/debug"
)

func (s *Service) handleRepoError(err error, operation string, ctx ...interface{}) error {
	logCtx := []interface{}{"operation", operation, "error", err}
	logCtx = append(logCtx, ctx...)

	switch {
	case errors.Is(err, students_models.ErrDuplicateContract),
		errors.Is(err, students_models.ErrInvalidClientID),
		errors.Is(err, students_models.ErrStudentNotFound),
		errors.Is(err, students_models.ErrStudentAlreadyDeleted),
		errors.Is(err, students_models.ErrStudentNotDeleted),
		errors.Is(err, context.Canceled),
		errors.Is(err, context.DeadlineExceeded):
		s.logger.Warn("repository error", logCtx...)

	case errors.Is(err, students_models.ErrCreateFailed),
		errors.Is(err, students_models.ErrUpdateFailed),
		errors.Is(err, students_models.ErrDeleteFailed),
		errors.Is(err, students_models.ErrRestoreFailed):
		s.logger.Error("operation failed", logCtx...)

	default:
		s.logger.Error("unexpected repository error",
			append(logCtx, "stack", string(debug.Stack()))...)
	}

	return err
}

package groups_service

import (
	"backend/pkg/models/groups"
	"context"
	"errors"
	"runtime/debug"
)

func (s *Service) handleRepoError(err error, operation string, ctx ...interface{}) error {
	logCtx := []interface{}{"operation", operation, "error", err}
	logCtx = append(logCtx, ctx...)

	switch {
	case errors.Is(err, groups_models.ErrGroupNotFound),
		errors.Is(err, groups_models.ErrDuplicateGroupName),
		errors.Is(err, groups_models.ErrInvalidGroupID),
		errors.Is(err, groups_models.ErrInvalidClientID),
		errors.Is(err, groups_models.ErrInvalidCuratorID),
		errors.Is(err, groups_models.ErrGroupNameRequired),
		errors.Is(err, groups_models.ErrGroupNameTooLong),
		errors.Is(err, groups_models.ErrInvalidPageSize),
		errors.Is(err, groups_models.ErrForeignKeyViolation),
		errors.Is(err, context.Canceled),
		errors.Is(err, context.DeadlineExceeded):
		s.logger.Warn("repository error", logCtx...)

	case errors.Is(err, groups_models.ErrCreateFailed),
		errors.Is(err, groups_models.ErrGetFailed),
		errors.Is(err, groups_models.ErrListFailed),
		errors.Is(err, groups_models.ErrUpdateFailed),
		errors.Is(err, groups_models.ErrDeleteFailed):
		s.logger.Error("operation failed", logCtx...)

	default:
		s.logger.Error("unexpected repository error",
			append(logCtx, "stack", string(debug.Stack()))...)
	}

	return err
}

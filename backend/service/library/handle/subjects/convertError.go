package subjects_handle

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// В файле subjects_handle.go
func (s *serverAPI) convertError(err error) error {
	switch {
	// Ошибки валидации
	case errors.Is(err, library_models.ErrEmptyName):
		return status.Error(codes.InvalidArgument, "subject name cannot be empty")
	case errors.Is(err, library_models.ErrInvalidSubjectID):
		return status.Error(codes.InvalidArgument, "invalid subject ID")

	// Ошибки бизнес-логики
	case errors.Is(err, library_models.ErrNotFound):
		return status.Error(codes.NotFound, "subject not found")
	case errors.Is(err, library_models.ErrDuplicateName):
		return status.Error(codes.AlreadyExists, "subject name already exists")
	case errors.Is(err, library_models.ErrDeleteConflict):
		return status.Error(codes.FailedPrecondition, "subject cannot be deleted due to existing references")

	// Ошибки контекста
	case errors.Is(err, context.Canceled):
		return status.Error(codes.Canceled, "request canceled")
	case errors.Is(err, context.DeadlineExceeded):
		return status.Error(codes.DeadlineExceeded, "deadline exceeded")

	// Ошибки доступа
	case errors.Is(err, library_models.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "permission denied")

	// Все остальные ошибки считаем внутренними
	default:
		s.logger.Error("Internal subject service error", "error", err)
		return status.Error(codes.Internal, "internal error")
	}
}

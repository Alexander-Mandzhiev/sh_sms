package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

var (
	ErrSoftDeleted         = errors.New("resource is soft deleted")
	ErrAlreadyExists       = errors.New("permission already exists")
	ErrNotFound            = errors.New("permission not found")
	ErrInvalidArgument     = errors.New("invalid argument")
	ErrPermissionDenied    = errors.New("permission denied")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrInvalidUUID         = errors.New("invalid uuid")
	ErrInvalidState        = errors.New("invalid object state")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, "permission not found")
	case errors.Is(err, ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, "invalid request data")
	case errors.Is(err, ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "permission code already exists")
	case errors.Is(err, ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "operation not allowed")
	case errors.Is(err, ErrForeignKeyViolation):
		return status.Error(codes.FailedPrecondition, "permission is in use")
	case errors.Is(err, ErrInvalidUUID):
		return status.Error(codes.InvalidArgument, "invalid UUID format")
	case errors.Is(err, ErrSoftDeleted):
		return status.Error(codes.FailedPrecondition, "permission already deleted")
	default:
		s.logger.Error("unhandled permissions error", slog.String("error", err.Error()))
		return status.Error(codes.Internal, "internal server error")
	}
}

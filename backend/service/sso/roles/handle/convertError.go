package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

var (
	ErrInvalidArgument       = errors.New("invalid argument")
	ErrPermissionDenied      = errors.New("permission denied")
	ErrAlreadyExists         = errors.New("role already exists")
	ErrNotFound              = errors.New("role not found")
	ErrUnauthenticated       = errors.New("unauthenticated")
	ErrRoleHierarchyConflict = errors.New("role hierarchy conflict")
	ErrRoleInUse             = errors.New("role is in use")
	ErrInvalidRoleLevel      = errors.New("invalid role level")
	ErrParentRoleNotFound    = errors.New("parent role not found")
	ErrPermissionNotFound    = errors.New("permission not found")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, "role not found")
	case errors.Is(err, ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, "invalid request data")
	case errors.Is(err, ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "role already exists")
	case errors.Is(err, ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "permission denied")
	case errors.Is(err, ErrUnauthenticated):
		return status.Error(codes.Unauthenticated, "authentication required")
	case errors.Is(err, ErrRoleHierarchyConflict):
		return status.Error(codes.FailedPrecondition, "role hierarchy conflict")
	case errors.Is(err, ErrRoleInUse):
		return status.Error(codes.FailedPrecondition, "role is assigned to users")
	case errors.Is(err, ErrInvalidRoleLevel):
		return status.Error(codes.InvalidArgument, "invalid role level")
	case errors.Is(err, ErrParentRoleNotFound):
		return status.Error(codes.NotFound, "parent role not found")
	case errors.Is(err, ErrPermissionNotFound):
		return status.Error(codes.NotFound, "permission not found")
	default:
		s.logger.Error("unhandled error", slog.Any("error", err))
		return status.Error(codes.Internal, "internal server error")
	}
}

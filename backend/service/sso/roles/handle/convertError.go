package handle

import (
	"backend/service/constants"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, constants.ErrNotFound):
		return status.Error(codes.NotFound, "role not found")
	case errors.Is(err, constants.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, "invalid request data")
	case errors.Is(err, constants.ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "role already exists")
	case errors.Is(err, constants.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "permission denied")
	case errors.Is(err, constants.ErrUnauthenticated):
		return status.Error(codes.Unauthenticated, "authentication required")
	case errors.Is(err, constants.ErrRoleHierarchyConflict):
		return status.Error(codes.FailedPrecondition, "role hierarchy conflict")
	case errors.Is(err, constants.ErrRoleInUse):
		return status.Error(codes.FailedPrecondition, "role is assigned to users")
	case errors.Is(err, constants.ErrInvalidRoleLevel):
		return status.Error(codes.InvalidArgument, "invalid role level")
	case errors.Is(err, constants.ErrParentRoleNotFound):
		return status.Error(codes.NotFound, "parent role not found")
	case errors.Is(err, constants.ErrPermissionNotFound):
		return status.Error(codes.NotFound, "permission not found")
	default:
		s.logger.Error("unhandled error", slog.Any("error", err))
		return status.Error(codes.Internal, "internal server error")
	}
}

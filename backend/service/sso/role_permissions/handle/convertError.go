package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrPermissionNotFound = errors.New("permission not found")
	ErrRoleNotFound       = errors.New("role not found")
	ErrInactiveEntity     = errors.New("inactive entity")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, "invalid argument")
	case errors.Is(err, ErrRoleNotFound):
		return status.Error(codes.NotFound, "role not found")
	case errors.Is(err, ErrPermissionNotFound):
		return status.Error(codes.NotFound, "permission not found")
	case errors.Is(err, ErrInactiveEntity):
		return status.Error(codes.FailedPrecondition, "entity is inactive")
	default:
		return status.Error(codes.Internal, "internal error")
	}
}

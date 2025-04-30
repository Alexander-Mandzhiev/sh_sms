package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrRoleNotFound      = errors.New("role not found")
	ErrAssignmentExists  = errors.New("role assignment already exists")
	ErrInvalidExpiration = errors.New("invalid expiration time")
	ErrInvalidArgument   = errors.New("invalid argument")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrUserNotFound):
		return status.Error(codes.NotFound, "user not found")
	case errors.Is(err, ErrRoleNotFound):
		return status.Error(codes.NotFound, "role not found")
	case errors.Is(err, ErrAssignmentExists):
		return status.Error(codes.AlreadyExists, "role assignment exists")
	case errors.Is(err, ErrInvalidExpiration):
		return status.Error(codes.InvalidArgument, "invalid expiration time")
	default:
		return status.Error(codes.Internal, "internal error")
	}
}

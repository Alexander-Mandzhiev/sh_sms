package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, "resource not found")
	case errors.Is(err, ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, "invalid request")
	case errors.Is(err, ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "resource already exists")
	case errors.Is(err, ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "access denied")
	case errors.Is(err, ErrUnauthenticated):
		return status.Error(codes.Unauthenticated, "authentication required")
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}

package handle

import (
	"backend/service/constants"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, constants.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, constants.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, constants.ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, constants.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, err.Error())
	default:
		return status.Error(codes.Internal, constants.ErrInternal.Error())
	}
}

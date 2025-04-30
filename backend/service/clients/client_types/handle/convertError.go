package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrClientTypeNotFound = errors.New("client type not found")
	ErrInvalidCode        = errors.New("invalid client type code")
	ErrCodeAlreadyExists  = errors.New("code already exists")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrDeletionConflict   = errors.New("cannot delete used client type")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrClientTypeNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, ErrInvalidCode),
		errors.Is(err, ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrCodeAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, ErrDeletionConflict):
		return status.Error(codes.FailedPrecondition, err.Error())
	default:
		s.logger.Error("internal error", "error", err)
		return status.Error(codes.Internal, "internal server error")
	}
}

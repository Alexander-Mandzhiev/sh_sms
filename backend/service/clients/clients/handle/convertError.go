package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound         = errors.New("client not found")
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrDeletionConflict = errors.New("cannot delete used client")
	ErrInternal         = errors.New("internal server error")
	ErrCodeExists       = errors.New("code exists")
	ErrConflict         = errors.New("conflict")
	ErrAlreadyActive    = errors.New("already active")
	ErrTimeout          = errors.New("request timeout")
	ErrDatabase         = errors.New("database error")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, "client not found")
	case errors.Is(err, ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrDeletionConflict):
		return status.Error(codes.FailedPrecondition, err.Error())
	case errors.Is(err, ErrCodeExists), errors.Is(err, ErrConflict):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, ErrAlreadyActive):
		return status.Error(codes.FailedPrecondition, err.Error())
	case errors.Is(err, ErrTimeout):
		return status.Error(codes.DeadlineExceeded, err.Error())
	case errors.Is(err, ErrDatabase):
		return status.Error(codes.Unavailable, "service unavailable")
	default:
		s.logger.Error("internal error", "error", err)
		return status.Error(codes.Internal, "internal server error")
	}
}

package classes_handle

import (
	sl "backend/pkg/logger"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidGrade  = errors.New("invalid grade")
	ErrClassNotFound = errors.New("class not found")
	ErrInvalidId     = errors.New("invalid id")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrInvalidId):
		return status.Error(codes.InvalidArgument, "invalid id")
	case errors.Is(err, ErrClassNotFound):
		return status.Error(codes.NotFound, "class not found")
	case errors.Is(err, ErrInvalidGrade):
		return status.Error(codes.InvalidArgument, "invalid class grade")
	default:
		s.logger.Error("Internal error", sl.Err(err, true))
		return status.Error(codes.Internal, "internal error")
	}
}

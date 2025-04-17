package service

import (
	"backend/service/apps/constants"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"reflect"
)

func (s *Service) convertError(err error) error {
	switch {
	case errors.Is(err, constants.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, constants.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, constants.ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, constants.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, constants.ErrInternal):
		return status.Error(codes.Internal, "internal server error")
	default:
		s.logger.Error("unhandled error type", slog.Any("error", err), slog.String("error_type", reflect.TypeOf(err).String()))
		return status.Error(codes.Internal, "internal server error")
	}
}

package handle

import (
	"backend/service/constants"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, constants.ErrInvalidID):
		return status.Error(codes.InvalidArgument, "invalid id")
	case errors.Is(err, constants.ErrEmptyCode):
		return status.Error(codes.InvalidArgument, "code cannot be empty")
	case errors.Is(err, constants.ErrConflictParams):
		return status.Error(codes.InvalidArgument, "conflicting parameters")
	case errors.Is(err, constants.ErrNoUpdateFields):
		return status.Error(codes.InvalidArgument, "no fields to update")
	case errors.Is(err, constants.ErrInvalidPagination):
		return status.Error(codes.InvalidArgument, "invalid pagination parameters")
	case errors.Is(err, constants.ErrIdentifierRequired):
		return status.Error(codes.InvalidArgument, "either id or code must be provided")
	case errors.Is(err, constants.ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "application already exists")
	case errors.Is(err, constants.ErrNotFound):
		return status.Error(codes.NotFound, "application not found")
	case errors.Is(err, constants.ErrEmptyName):
		return status.Error(codes.InvalidArgument, "name cannot be empty")
	case errors.Is(err, constants.ErrUpdateConflict):
		return status.Error(codes.Aborted, "update conflict")
	case errors.Is(err, constants.ErrVersionConflict):
		return status.Error(codes.Aborted, "version conflict")

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}

package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidID          = errors.New("invalid id")
	ErrEmptyName          = errors.New("name cannot be empty")
	ErrEmptyCode          = errors.New("code cannot be empty")
	ErrInvalidName        = errors.New("name length invalid ")
	ErrInvalidCode        = errors.New("code length invalid ")
	ErrConflictParams     = errors.New("conflicting parameters")
	ErrNoUpdateFields     = errors.New("no fields to update")
	ErrInvalidPagination  = errors.New("invalid pagination parameters")
	ErrIdentifierRequired = errors.New("either id or code must be provided")
	ErrAlreadyExists      = errors.New("application already exists")
	ErrNotFound           = errors.New("application not found")
	ErrUpdateConflict     = errors.New("update conflict")
	ErrVersionConflict    = errors.New("version conflict")
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, ErrInvalidID):
		return status.Error(codes.InvalidArgument, "invalid id")
	case errors.Is(err, ErrEmptyCode):
		return status.Error(codes.InvalidArgument, "code cannot be empty")
	case errors.Is(err, ErrConflictParams):
		return status.Error(codes.InvalidArgument, "conflicting parameters")
	case errors.Is(err, ErrNoUpdateFields):
		return status.Error(codes.InvalidArgument, "no fields to update")
	case errors.Is(err, ErrInvalidPagination):
		return status.Error(codes.InvalidArgument, "invalid pagination parameters")
	case errors.Is(err, ErrIdentifierRequired):
		return status.Error(codes.InvalidArgument, "either id or code must be provided")
	case errors.Is(err, ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "application already exists")
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, "application not found")
	case errors.Is(err, ErrEmptyName):
		return status.Error(codes.InvalidArgument, "name cannot be empty")
	case errors.Is(err, ErrUpdateConflict):
		return status.Error(codes.Aborted, "update conflict")
	case errors.Is(err, ErrVersionConflict):
		return status.Error(codes.Aborted, "version conflict")

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}

package handle

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidID          = errors.New("invalid id")
	ErrEmptyCode          = errors.New("code cannot be empty")
	ErrConflictParams     = errors.New("conflicting parameters")
	ErrNoUpdateFields     = errors.New("no fields to update")
	ErrInvalidPagination  = errors.New("invalid pagination parameters")
	ErrIdentifierRequired = errors.New("either id or code must be provided")
	ErrAlreadyExists      = errors.New("application already exists")
	ErrNotFound           = errors.New("application not found")
)

func convertError(err error) error {
	switch {
	case errors.Is(err, ErrInvalidID),
		errors.Is(err, ErrEmptyCode),
		errors.Is(err, ErrInvalidPagination),
		errors.Is(err, ErrConflictParams),
		errors.Is(err, ErrNoUpdateFields),
		errors.Is(err, ErrIdentifierRequired):
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "application already exists")

	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, "application not found")

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}

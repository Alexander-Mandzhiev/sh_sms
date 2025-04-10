package app_manager_handle

import (
	apps "backend/service/clients/service/apps_manager"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertError(err error) error {
	switch {
	case errors.Is(err, apps.ErrInvalidID):
		return status.Error(codes.InvalidArgument, "invalid application ID format")
	case errors.Is(err, apps.ErrEmptyCode):
		return status.Error(codes.InvalidArgument, "application code cannot be empty")
	case errors.Is(err, apps.ErrInvalidPagination):
		return status.Error(codes.InvalidArgument, "invalid pagination parameters")
	case errors.Is(err, apps.ErrConflictParams):
		return status.Error(codes.InvalidArgument, "conflicting identifier parameters")
	case errors.Is(err, apps.ErrNoUpdateFields):
		return status.Error(codes.InvalidArgument, "no fields provided for update")
	case errors.Is(err, apps.ErrIdentifierRequired):
		return status.Error(codes.InvalidArgument, "either ID or code must be provided")
	case errors.Is(err, apps.ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "application with this code already exists")
	case errors.Is(err, apps.ErrNotFound):
		return status.Error(codes.NotFound, "application not found")
	case errors.Is(err, apps.ErrInvalidName):
		return status.Error(codes.InvalidArgument, "invalid application name format")
	case errors.Is(err, apps.ErrMaxCountExceeded):
		return status.Error(codes.InvalidArgument, "maximum items per page exceeded")
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}

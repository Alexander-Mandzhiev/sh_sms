package groups_handle

import (
	groups_models "backend/pkg/models/groups"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, context.Canceled):
		return status.Error(codes.Canceled, "request canceled")
	case errors.Is(err, context.DeadlineExceeded):
		return status.Error(codes.DeadlineExceeded, "deadline exceeded")
	case errors.Is(err, groups_models.ErrGroupNameRequired):
		return status.Error(codes.InvalidArgument, "group name is required")
	case errors.Is(err, groups_models.ErrGroupNameTooLong):
		return status.Error(codes.InvalidArgument, "group name exceeds maximum length (100 characters)")
	case errors.Is(err, groups_models.ErrInvalidPageSize):
		return status.Error(codes.InvalidArgument, "page size must be between 1 and 100")
	case errors.Is(err, groups_models.ErrInvalidGroupID):
		return status.Error(codes.InvalidArgument, "invalid group ID format")
	case errors.Is(err, groups_models.ErrInvalidClientID):
		return status.Error(codes.InvalidArgument, "invalid client ID format")
	case errors.Is(err, groups_models.ErrInvalidCuratorID):
		return status.Error(codes.InvalidArgument, "invalid curator ID format")
	case errors.Is(err, groups_models.ErrFilterValueTooLong):
		return status.Error(codes.InvalidArgument, "filter value exceeds maximum length")
	case errors.Is(err, groups_models.ErrInvalidCursor):
		return status.Error(codes.InvalidArgument, "invalid cursor value")
	case errors.Is(err, groups_models.ErrDuplicateGroupName):
		return status.Error(codes.AlreadyExists, "group name already exists for this client")
	case errors.Is(err, groups_models.ErrGroupNotFound):
		return status.Error(codes.NotFound, "group not found")
	case errors.Is(err, groups_models.ErrForeignKeyViolation):
		return status.Error(codes.FailedPrecondition, "cannot perform operation due to existing dependencies")
	case errors.Is(err, groups_models.ErrCreateFailed):
		return status.Error(codes.Internal, "failed to create group")
	case errors.Is(err, groups_models.ErrGetFailed):
		return status.Error(codes.Internal, "failed to get group")
	case errors.Is(err, groups_models.ErrListFailed):
		return status.Error(codes.Internal, "failed to list groups")
	case errors.Is(err, groups_models.ErrUpdateFailed):
		return status.Error(codes.Internal, "failed to update group")
	case errors.Is(err, groups_models.ErrDeleteFailed):
		return status.Error(codes.Internal, "failed to delete group")
	default:
		s.logger.Error("Internal server error", slog.String("error", err.Error()), slog.Any("type", "unhandled_error"))
		return status.Error(codes.Internal, "internal server error")
	}
}

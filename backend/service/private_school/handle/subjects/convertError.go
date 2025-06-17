package subjects_handle

import (
	subjects_models "backend/pkg/models/subject"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, subjects_models.ErrEmptySubjectName):
		return status.Error(codes.InvalidArgument, "subject name cannot be empty")
	case errors.Is(err, subjects_models.ErrInvalidSubjectID):
		return status.Error(codes.InvalidArgument, "invalid subject ID")
	case errors.Is(err, subjects_models.ErrNotFoundSubjectName):
		return status.Error(codes.NotFound, "subject not found")
	case errors.Is(err, subjects_models.ErrDuplicateSubjectName):
		return status.Error(codes.AlreadyExists, "subject name already exists")
	case errors.Is(err, subjects_models.ErrDeleteSubjectConflict):
		return status.Error(codes.FailedPrecondition, "subject cannot be deleted due to existing references")
	case errors.Is(err, context.Canceled):
		return status.Error(codes.Canceled, "request canceled")
	case errors.Is(err, context.DeadlineExceeded):
		return status.Error(codes.DeadlineExceeded, "deadline exceeded")
	case errors.Is(err, subjects_models.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "permission denied")
	default:
		s.logger.Error("Internal subject service error", "error", err)
		return status.Error(codes.Internal, "internal error")
	}
}

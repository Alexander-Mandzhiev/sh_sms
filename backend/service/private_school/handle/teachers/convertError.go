package teachers_handle

import (
	"backend/pkg/models/private_school"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, private_school_models.ErrEmptyFullName):
		return status.Error(codes.InvalidArgument, "full name cannot be empty")
	case errors.Is(err, private_school_models.ErrInvalidPhone):
		return status.Error(codes.InvalidArgument, "invalid phone format")
	case errors.Is(err, private_school_models.ErrInvalidEmail):
		return status.Error(codes.InvalidArgument, "invalid email format")
	case errors.Is(err, private_school_models.ErrInvalidTeacherID):
		return status.Error(codes.InvalidArgument, "invalid teacher ID")
	case errors.Is(err, private_school_models.ErrTeacherNotFound):
		return status.Error(codes.NotFound, "teacher not found")
	case errors.Is(err, private_school_models.ErrDuplicateTeacher):
		return status.Error(codes.AlreadyExists, "teacher already exists")
	case errors.Is(err, private_school_models.ErrInvalidClient):
		return status.Error(codes.InvalidArgument, "invalid client reference")
	case errors.Is(err, private_school_models.ErrCreateFailed):
		return status.Error(codes.Internal, "failed to create teacher")
	case errors.Is(err, private_school_models.ErrDeleteTeacherConflict):
		return status.Error(codes.FailedPrecondition, "teacher has active references")
	case errors.Is(err, private_school_models.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "permission denied")
	case errors.Is(err, context.Canceled):
		return status.Error(codes.Canceled, "request canceled")
	case errors.Is(err, context.DeadlineExceeded):
		return status.Error(codes.DeadlineExceeded, "deadline exceeded")

	default:
		s.logger.Error("Internal teacher service error", slog.Any("error", err), slog.String("error_type", fmt.Sprintf("%T", err)))
		return status.Error(codes.Internal, "internal error: "+err.Error())
	}
}

package students_handle

import (
	"backend/pkg/models/students"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, students_models.ErrEmptyFullName):
		return status.Error(codes.InvalidArgument, "full name is required")
	case errors.Is(err, students_models.ErrEmptyContractNumber):
		return status.Error(codes.InvalidArgument, "contract number is required")
	case errors.Is(err, students_models.ErrEmptyPhone):
		return status.Error(codes.InvalidArgument, "phone is required")
	case errors.Is(err, students_models.ErrEmptyEmail):
		return status.Error(codes.InvalidArgument, "email is required")
	case errors.Is(err, students_models.ErrInvalidPhone):
		return status.Error(codes.InvalidArgument, "invalid phone format")
	case errors.Is(err, students_models.ErrInvalidEmail):
		return status.Error(codes.InvalidArgument, "invalid email format")
	case errors.Is(err, students_models.ErrFullNameTooLong):
		return status.Error(codes.InvalidArgument, "full name too long")
	case errors.Is(err, students_models.ErrContractNumberTooLong):
		return status.Error(codes.InvalidArgument, "contract number too long")
	case errors.Is(err, students_models.ErrEmailTooLong):
		return status.Error(codes.InvalidArgument, "email too long")
	case errors.Is(err, students_models.ErrPhoneTooShort):
		return status.Error(codes.InvalidArgument, "phone number too short")
	case errors.Is(err, students_models.ErrPhoneTooLong):
		return status.Error(codes.InvalidArgument, "phone number too long")
	case errors.Is(err, students_models.ErrDuplicateContract):
		return status.Error(codes.AlreadyExists, "contract number already exists")
	case errors.Is(err, students_models.ErrInvalidClientID):
		return status.Error(codes.InvalidArgument, "invalid client ID")
	case errors.Is(err, students_models.ErrInvalidStudentID):
		return status.Error(codes.InvalidArgument, "invalid student ID format")
	case errors.Is(err, students_models.ErrStudentNotFound):
		return status.Error(codes.NotFound, "student not found")
	case errors.Is(err, students_models.ErrStudentAlreadyDeleted):
		return status.Error(codes.FailedPrecondition, "student already deleted")
	case errors.Is(err, students_models.ErrStudentNotDeleted):
		return status.Error(codes.FailedPrecondition, "student is not deleted")
	case errors.Is(err, students_models.ErrCreateFailed):
		return status.Error(codes.Internal, "student creation failed")
	case errors.Is(err, students_models.ErrInvalidCursor):
		return status.Error(codes.InvalidArgument, "invalid cursor format")
	default:
		s.logger.Error("Internal server error", slog.String("error", err.Error()))
		return status.Error(codes.Internal, "internal server error")
	}
}

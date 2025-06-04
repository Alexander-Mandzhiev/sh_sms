package attachment_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) convertError(err error) error {
	if se, ok := status.FromError(err); ok {
		return se.Err()
	}
	switch {
	case errors.Is(err, library_models.ErrAttachmentNotFound):
		return status.Error(codes.NotFound, "attachment not found")
	case errors.Is(err, library_models.ErrAttachmentAlreadyExists):
		return status.Error(codes.AlreadyExists, "attachment already exists")
	case errors.Is(err, library_models.ErrAttachmentExistsButDeleted):
		return status.Error(codes.FailedPrecondition, "attachment exists but is deleted")
	case errors.Is(err, library_models.ErrAttachmentAlreadyActive):
		return status.Error(codes.FailedPrecondition, "attachment is already active")
	case errors.Is(err, library_models.ErrEmptyFileURL):
		return status.Error(codes.InvalidArgument, "file URL cannot be empty")
	case errors.Is(err, library_models.ErrAttachmentRestoreConflict):
		return status.Error(codes.Aborted, "active attachment already exists, cannot restore")
	case errors.Is(err, library_models.ErrAttachmentUpdateConflict):
		return status.Error(codes.Aborted, "conflict during attachment update")
	case errors.Is(err, sql.ErrNoRows):
		return status.Error(codes.NotFound, "no attachments found")
	case errors.Is(err, library_models.ErrEmptyFileURL):
		return status.Error(codes.InvalidArgument, "file URL cannot be empty")
	case isGRPCStatusError(err):
		return err
	default:
		s.logger.Error("Internal attachment error", slog.String("error", err.Error()), sl.Err(err, true))
		return status.Error(codes.Internal, "internal server error")
	}
}

func isGRPCStatusError(err error) bool {
	_, ok := status.FromError(err)
	return ok
}

package books_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) convertError(err error) error {
	switch {
	case errors.Is(err, library_models.ErrInvalidID):
		return status.Error(codes.InvalidArgument, "invalid book id")
	case errors.Is(err, library_models.ErrNotFound):
		return status.Error(codes.NotFound, "book not found")
	case errors.Is(err, library_models.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "permission denied")
	case errors.Is(err, library_models.ErrBookInvalidClientID):
		return status.Error(codes.InvalidArgument, "invalid client id format")
	case errors.Is(err, library_models.ErrBookEmptyTitle):
		return status.Error(codes.InvalidArgument, "title cannot be empty")
	case errors.Is(err, library_models.ErrBookEmptyAuthor):
		return status.Error(codes.InvalidArgument, "author cannot be empty")
	case errors.Is(err, library_models.ErrBookInvalidSubjectID):
		return status.Error(codes.InvalidArgument, "invalid subject id")
	case errors.Is(err, library_models.ErrBookInvalidClassID):
		return status.Error(codes.InvalidArgument, "invalid class id")
	case errors.Is(err, library_models.ErrBookDescriptionLong):
		return status.Error(codes.InvalidArgument, "description too long")
	case errors.Is(err, library_models.ErrInvalidPageSize):
		return status.Error(codes.InvalidArgument, "invalid page size")
	case errors.Is(err, library_models.ErrClientIDRequired):
		return status.Error(codes.InvalidArgument, "client ID is required")
	default:
		s.logger.Error("Internal error", sl.Err(err, true))
		return status.Error(codes.Internal, "internal error")
	}
}

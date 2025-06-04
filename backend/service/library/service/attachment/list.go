package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) ListAttachmentsByBook(ctx context.Context, bookID int64, includeDeleted bool) ([]*library_models.Attachment, error) {
	const op = "service.Library.Attachments.ListByBook"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", bookID), slog.Bool("include_deleted", includeDeleted))
	logger.Debug("Listing attachments")

	if err := library_models.ValidateBookID(bookID); err != nil {
		logger.Warn("Invalid book ID")
		return nil, status.Error(codes.InvalidArgument, "invalid book ID")
	}

	attachments, err := s.provider.ListByBook(ctx, bookID, includeDeleted)
	if err != nil {
		logger.Error("Failed to list attachments", "error", err)
		return nil, err
	}

	logger.Debug("Attachments listed successfully", slog.Int("count", len(attachments)))
	return attachments, nil
}

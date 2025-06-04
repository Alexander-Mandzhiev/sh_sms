package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) DeleteAttachment(ctx context.Context, bookID int64, format string) error {
	const op = "service.Library.Attachments.Delete"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", bookID), slog.String("format", format))
	logger.Debug("Delete attachment")

	if err := library_models.ValidateBookID(bookID); err != nil {
		logger.Warn("Invalid book ID", "error", err)
		return err
	}
	if err := library_models.ValidateAttachmentFormat(format); err != nil {
		logger.Warn("Invalid format", "error", err)
		return err
	}

	if err := s.provider.DeleteAttachment(ctx, bookID, format); err != nil {
		logger.Error("Failed to deleting attachment", "error", err)
		return err
	}

	logger.Debug("Attachment retrieved successfully")
	return nil
}

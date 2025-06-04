package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) GetAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error) {
	const op = "service.Library.Attachments.Get"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", bookID), slog.String("format", format))
	logger.Debug("Getting attachment")

	if err := library_models.ValidateBookID(bookID); err != nil {
		logger.Warn("Invalid book ID", "error", err)
		return nil, err
	}
	if err := library_models.ValidateAttachmentFormat(format); err != nil {
		logger.Warn("Invalid format", "error", err)
		return nil, err
	}

	attachment, err := s.provider.GetAttachment(ctx, bookID, format)
	if err != nil {
		logger.Error("Failed to get attachment", "error", err)
		return nil, library_models.ErrAttachmentRestoreConflict
	}

	logger.Debug("Attachment retrieved successfully")
	return attachment, nil
}

package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) RestoreAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error) {
	const op = "service.Library.Attachments.Restore"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", bookID), slog.String("format", format))
	logger.Debug("Restore attachment")

	if err := library_models.ValidateBookID(bookID); err != nil {
		logger.Warn("Invalid book ID", "error", err)
		return nil, err
	}
	if err := library_models.ValidateAttachmentFormat(format); err != nil {
		logger.Warn("Invalid format", "error", err)
		return nil, err
	}

	attachment, err := s.provider.RestoreAttachment(ctx, bookID, format)
	if err != nil {
		logger.Error("Failed to restore attachment", "error", err)
		return nil, err
	}

	logger.Debug("Attachment restore successfully")
	return attachment, nil
}

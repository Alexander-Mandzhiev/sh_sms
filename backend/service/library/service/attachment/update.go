package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) UpdateAttachment(ctx context.Context, params *library_models.UpdateAttachmentRequest) (*library_models.Attachment, error) {
	const op = "service.Library.Attachments.Update"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", params.BookID), slog.String("format", params.Format))
	logger.Debug("Updating attachment")

	if err := params.Validate(); err != nil {
		logger.Error("Validation failed", "error", err)
		return nil, err
	}

	currentAttachment, err := s.provider.GetAttachment(ctx, params.BookID, params.Format)
	if err != nil {
		logger.Error("Failed to get attachment", "error", err)
		return nil, library_models.ErrAttachmentRestoreConflict
	}

	currentAttachment.FileURL = params.NewFileURL

	if err = s.provider.UpdateAttachment(ctx, currentAttachment); err != nil {
		logger.Error("Failed to update attachment", "error", err)
		return nil, err
	}

	logger.Info("Attachment updated successfully")
	return currentAttachment, nil
}

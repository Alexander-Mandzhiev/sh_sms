package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) CreateAttachment(ctx context.Context, params *library_models.CreateAttachmentRequest) (*library_models.Attachment, error) {
	const op = "service.Library.Attachments.Create"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", params.BookId), slog.String("format", params.Format))
	logger.Debug("Creating attachment")

	if err := params.Validate(); err != nil {
		logger.Error("Validation failed", "error", err)
		return nil, err
	}

	attachment := &library_models.Attachment{
		BookID: params.BookId,
		Format: params.Format,
		FileID: params.FileId,
	}

	if err := s.provider.CreateAttachment(ctx, attachment); err != nil {
		logger.Error("Failed to create attachment", "error", err)
		return nil, err
	}

	logger.Info("Attachment created successfully")
	return attachment, nil
}

package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) DeleteAttachment(ctx context.Context, fileId string) error {
	const op = "service.Library.Attachments.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("file_id", fileId))
	logger.Debug("Delete attachment")

	if err := library_models.ValidateFileID(fileId); err != nil {
		logger.Warn("Invalid file ID", "error", err)
		return err
	}

	if err := s.provider.DeleteAttachment(ctx, fileId); err != nil {
		logger.Error("Failed to deleting attachment", "error", err)
		return err
	}

	logger.Debug("Attachment retrieved successfully")
	return nil
}

package attachments_service

import (
	library "backend/protos/gen/go/library"
	"context"
	"fmt"
	"log/slog"
)

func (s *AttachmentsService) DeleteAttachment(ctx context.Context, fileId string) error {
	const op = "service.Gateway.Attachments.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("file_id", fileId))

	if _, err := s.client.DeleteAttachment(ctx, &library.DeleteAttachmentRequest{FileId: fileId}); err != nil {
		logger.Error("Failed to delete attachment record", "error", err)
		return fmt.Errorf("failed to delete record: %w", err)
	}

	if err := s.storage.DeleteFile(ctx, fileId); err != nil {
		logger.Error("Failed to delete file from storage", "error", err)
		return fmt.Errorf("failed to delete file: %w", err)
	}

	logger.Info("Attachment and file deleted successfully")
	return nil
}

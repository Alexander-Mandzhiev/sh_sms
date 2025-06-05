// attachments_service.go
package attachments_service

import (
	library "backend/protos/gen/go/library"
	"context"
	"fmt"
	"io"
	"log/slog"
)

func (s *AttachmentsService) GetAttachment(ctx context.Context, bookId int64, format string) (io.ReadSeekCloser, string, int64, error) {
	const op = "service.Attachments.GetAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", bookId), slog.String("format", format))
	logger.Debug("Getting attachment metadata")

	att, err := s.client.GetAttachment(ctx, &library.GetAttachmentRequest{BookId: bookId, Format: format})
	if err != nil {
		logger.Error("Failed to get attachment metadata", "error", err, "details", "gRPC call failed")
		return nil, "", 0, fmt.Errorf("failed to get attachment metadata: %w", err)
	}

	logger.Debug("Attachment metadata retrieved", "file_id", att.FileId)

	file, filePath, size, err := s.storage.GetFile(ctx, att.FileId)
	if err != nil {
		logger.Error("Failed to get file from storage", "file_id", att.FileId, "error", err, "details", "storage operation failed")
		return nil, "", 0, fmt.Errorf("failed to get file from storage: %w", err)
	}

	logger.Info("Attachment and file retrieved successfully", "file_size", fmt.Sprintf("%d bytes", size), "file_path", filePath)
	return file, filePath, size, nil
}

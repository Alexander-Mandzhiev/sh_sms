package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"strings"
)

func (s *AttachmentsService) CreateAttachment(ctx context.Context, meta library_models.FileMetadata, file multipart.File) (*library_models.Attachment, error) {
	const op = "service.Gateway.Attachments.Create"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", meta.BookID), slog.String("format", meta.Format))

	uploadedFile, err := s.storage.SaveFile(ctx, meta, file)
	if err != nil {
		logger.Error("Failed to save file", "error", err)
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	createReq := &library_models.CreateAttachmentRequest{
		BookId: meta.BookID,
		Format: strings.ToUpper(meta.Format),
		FileId: uploadedFile.FilePath,
	}

	protoReq, err := createReq.CreateAttachmentRequestToProto()
	if err != nil {
		logger.Error("Failed to create protobuf request", "error", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	attachment, err := s.client.CreateAttachment(ctx, protoReq)
	if err != nil {
		if delErr := s.storage.DeleteFile(ctx, uploadedFile.FilePath); delErr != nil {
			logger.Error("Compensation failed: could not delete file", "file_path", uploadedFile.FilePath, "error", delErr)
		}
		logger.Error("Failed to create attachment record", "error", err)
		return nil, fmt.Errorf("failed to create attachment: %w", err)
	}
	logger.Info("Attachment created successfully")
	return library_models.AttachmentFromProto(attachment), nil
}

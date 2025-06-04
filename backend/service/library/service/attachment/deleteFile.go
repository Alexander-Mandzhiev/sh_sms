package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) DeleteFile(ctx context.Context, fileURL string) error {
	const op = "service.Library.Attachments.DeleteFile"
	logger := s.logger.With(slog.String("op", op), slog.String("file_url", fileURL))

	if fileURL == "" {
		logger.Warn("Empty file URL provided")
		return library_models.ErrEmptyFileURL
	}

	if err := s.fileStorage.DeleteFile(ctx, fileURL); err != nil {
		logger.Error("Failed to delete file", "error", err)
		return err
	}

	logger.Info("File deleted successfully")
	return nil
}

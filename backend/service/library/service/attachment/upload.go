package attachments_service

import (
	"backend/pkg/models/library"
	"context"
	"io"
)

func (s *Service) UploadFile(ctx context.Context, metadata *library_models.FileMetadata, file io.Reader) (*library_models.UploadedFile, error) {
	if metadata.BookID <= 0 {
		return nil, library_models.ErrInvalidID
	}
	if metadata.Format == "" {
		return nil, library_models.ErrInvalidAttachmentFormat
	}

	uploadedFile, err := s.fileStorage.SaveFile(ctx, *metadata, file)
	if err != nil {
		s.logger.Error("Failed to save file", "error", err)
		return nil, err
	}

	return &uploadedFile, nil
}

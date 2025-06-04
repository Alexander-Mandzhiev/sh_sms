package attachment_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) RestoreAttachment(ctx context.Context, req *library.RestoreAttachmentRequest) (*library.Attachment, error) {
	const op = "grpc.Attachment.RestoreAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", req.BookId), slog.String("format", req.GetFormat()))
	logger.Debug("Restore attachment request received")

	// Валидация параметров
	if err := library_models.ValidateBookID(req.BookId); err != nil {
		logger.Warn("Book ID validation failed", slog.String("error", err.Error()), sl.Err(err, false))
		return nil, err
	}
	if err := library_models.ValidateAttachmentFormat(req.Format); err != nil {
		logger.Warn("Format validation failed", slog.String("error", err.Error()), sl.Err(err, false))
		return nil, err
	}

	attachment, err := s.service.RestoreAttachment(ctx, req.BookId, req.Format)
	if err != nil {
		logger.Error("Failed to restore attachment", slog.String("error", err.Error()), sl.Err(err, true))
		return nil, s.convertError(err)
	}

	logger.Info("Attachment successfully restored")
	return attachment.AttachmentToProto(), nil
}

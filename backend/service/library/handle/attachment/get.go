package attachment_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) GetAttachment(ctx context.Context, req *library.GetAttachmentRequest) (*library.Attachment, error) {
	const op = "grpc.Attachment.GetAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", req.BookId), slog.String("format", req.GetFormat()))
	logger.Debug("get attachment request received")

	if err := library_models.ValidateBookID(req.BookId); err != nil {
		logger.Warn("book ID validation failed", sl.Err(err, true))
		return nil, err
	}
	if err := library_models.ValidateAttachmentFormat(req.Format); err != nil {
		logger.Warn("format validation failed", sl.Err(err, true))
		return nil, err
	}

	attachment, err := s.service.GetAttachment(ctx, req.BookId, req.Format)
	if err != nil {
		logger.Error("failed to get attachment", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	return attachment.AttachmentToProto(), nil
}

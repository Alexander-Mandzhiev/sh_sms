package attachment_handle

import (
	"context"
	"log/slog"

	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
)

func (s *serverAPI) UpdateAttachment(ctx context.Context, req *library.UpdateAttachmentRequest) (*library.Attachment, error) {
	const op = "grpc.Attachment.UpdateAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", req.BookId), slog.String("format", req.GetFormat()))
	logger.Debug("Update attachment request received")

	domainReq, err := library_models.UpdateAttachmentRequestFromProto(req)
	if err != nil {
		logger.Warn("Request validation failed", slog.String("error", err.Error()), sl.Err(err, false))
		return nil, err
	}

	updatedAttachment, err := s.service.UpdateAttachment(ctx, domainReq)
	if err != nil {
		logger.Error("Failed to update attachment", slog.String("error", err.Error()), sl.Err(err, true))
		return nil, s.convertError(err)
	}

	logger.Info("Attachment successfully updated")
	return updatedAttachment.AttachmentToProto(), nil
}

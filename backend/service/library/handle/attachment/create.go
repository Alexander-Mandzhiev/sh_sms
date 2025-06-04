package attachment_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateAttachment(ctx context.Context, req *library.CreateAttachmentRequest) (*library.Attachment, error) {
	const op = "grpc.Attachment.CreateAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", req.BookId), slog.String("format", req.GetFormat()))
	logger.Debug("create attachment request received")

	domainReq := library_models.CreateAttachmentRequestFromProto(req)
	attachment, err := s.service.CreateAttachment(ctx, domainReq)
	if err != nil {
		logger.Error("create attachment failed", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	return attachment.AttachmentToProto(), nil
}

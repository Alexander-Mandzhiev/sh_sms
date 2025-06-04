package attachment_handle

import (
	"context"
	"log/slog"

	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
)

func (s *serverAPI) ListAttachmentsByBook(ctx context.Context, req *library.ListAttachmentsByBookRequest) (*library.ListAttachmentsResponse, error) {
	const op = "grpc.Attachment.ListAttachmentsByBook"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", req.BookId))
	logger.Debug("List attachments by book request received")

	if err := library_models.ValidateBookID(req.BookId); err != nil {
		logger.Warn("Book ID validation failed", slog.String("error", err.Error()), sl.Err(err, false))
		return nil, err
	}

	attachments, err := s.service.ListAttachmentsByBook(ctx, req.BookId, req.IncludeDeleted)
	if err != nil {
		logger.Error("Failed to list attachments", slog.String("error", err.Error()), sl.Err(err, true))
		return nil, s.convertError(err)
	}

	pbAttachments := make([]*library.Attachment, 0, len(attachments))
	for _, a := range attachments {
		pbAttachments = append(pbAttachments, a.AttachmentToProto())
	}

	logger.Info("Attachments successfully listed", slog.Int("count", len(pbAttachments)), slog.Bool("include_deleted", req.IncludeDeleted))
	return &library.ListAttachmentsResponse{Attachments: pbAttachments}, nil
}

package attachment_handle

import (
	"context"
	"log/slog"

	sl "backend/pkg/logger"
	library_models "backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serverAPI) DeleteAttachment(ctx context.Context, req *library.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	const op = "grpc.Attachment.DeleteAttachment"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", req.BookId), slog.String("format", req.GetFormat()))
	logger.Debug("Delete attachment request received")

	if err := library_models.ValidateBookID(req.BookId); err != nil {
		logger.Warn("Book ID validation failed", slog.String("error", err.Error()), sl.Err(err, false))
		return nil, err
	}
	if err := library_models.ValidateAttachmentFormat(req.Format); err != nil {
		logger.Warn("Format validation failed", slog.String("error", err.Error()), sl.Err(err, false))
		return nil, err
	}

	if err := s.service.DeleteAttachment(ctx, req.BookId, req.Format); err != nil {
		logger.Error("Failed to delete attachment", slog.String("error", err.Error()), sl.Err(err, true))
		return nil, s.convertError(err)
	}

	logger.Info("Attachment successfully deleted")
	return &emptypb.Empty{}, nil
}

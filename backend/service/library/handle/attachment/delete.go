package attachment_handle

import (
	"context"
	"log/slog"

	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serverAPI) DeleteAttachment(ctx context.Context, req *library.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	const op = "grpc.Attachment.DeleteAttachment"
	logger := s.logger.With(slog.String("op", op), slog.String("file_id", req.FileId))
	logger.Debug("Delete attachment request received")

	if err := library_models.ValidateFileID(req.FileId); err != nil {
		logger.Warn("Format validation failed", slog.String("error", err.Error()), sl.Err(err, false))
		return nil, err
	}

	if err := s.service.DeleteAttachment(ctx, req.FileId); err != nil {
		logger.Error("Failed to delete attachment", slog.String("error", err.Error()), sl.Err(err, true))
		return nil, s.convertError(err)
	}

	logger.Info("Attachment successfully deleted")
	return &emptypb.Empty{}, nil
}

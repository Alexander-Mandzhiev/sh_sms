package attachment_handle

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	library "backend/protos/gen/go/library"
)

func (s *serverAPI) DeleteFile(ctx context.Context, req *library.DeleteFileRequest) (*emptypb.Empty, error) {
	const op = "grpc.Attachment.DeleteFile"
	logger := s.logger.With(slog.String("op", op), slog.String("file_url", req.GetFileUrl()))
	logger.Debug("Delete file request received")

	if req.GetFileUrl() == "" {
		logger.Warn("Empty file URL provided")
		return nil, status.Error(codes.InvalidArgument, "file_url is required")
	}

	if err := s.service.DeleteFile(ctx, req.GetFileUrl()); err != nil {
		logger.Error("Failed to delete file", "error", err)
		return nil, s.convertError(err)
	}

	logger.Info("File deleted successfully")
	return &emptypb.Empty{}, nil
}

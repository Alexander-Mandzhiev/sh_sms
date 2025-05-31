package file_format_handle

import (
	sl "backend/pkg/logger"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) FileFormatExists(ctx context.Context, req *library.FileFormatExistsRequest) (*library.FileFormatExistsResponse, error) {
	const op = "grpc.Library.FileFormats.FileFormatExists"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("FileFormatExists called", slog.String("format", req.GetFormat()))

	exists, err := s.service.FileFormatExists(ctx, req.GetFormat())
	if err != nil {
		logger.Error("Failed to check file format existence", sl.Err(err, true), slog.String("format", req.GetFormat()))
		return nil, s.convertError(err)
	}

	return &library.FileFormatExistsResponse{Exists: exists}, nil
}

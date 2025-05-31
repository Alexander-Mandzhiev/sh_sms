package file_format_handle

import (
	sl "backend/pkg/logger"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) ListFileFormats(ctx context.Context, _ *emptypb.Empty) (*library.ListFileFormatsResponse, error) {
	const op = "grpc.Library.FileFormats.ListFileFormats"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("ListFileFormats called")

	formats, err := s.service.ListFileFormats(ctx)
	if err != nil {
		logger.Error("Failed to list file formats", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	return &library.ListFileFormatsResponse{Formats: formats}, nil
}

package file_format_service

import (
	"context"
	"log/slog"
)

func (s *FileFormatService) ListFileFormats(ctx context.Context) ([]string, error) {
	const op = "service.Library.FileFormats.ListFileFormats"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("listing file formats")

	formats, err := s.provider.ListFileFormats(ctx)
	if err != nil {
		logger.Error("failed to list file formats", slog.String("error", err.Error()))
		return nil, err
	}

	logger.Debug("file formats listed successfully", slog.Int("count", len(formats)))
	return formats, nil
}

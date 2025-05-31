package file_format_service

import (
	"context"
	"log/slog"
)

func (s *FileFormatService) FileFormatExists(ctx context.Context, format string) (bool, error) {
	const op = "service.Library.FileFormats.FileFormatExists"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("checking file format existence", slog.String("format", format))

	exists, err := s.provider.FileFormatExists(ctx, format)
	if err != nil {
		logger.Error("failed to check file format existence", slog.String("error", err.Error()), slog.String("format", format))
		return false, err
	}

	logger.Debug("file format check completed", slog.String("format", format), slog.Bool("exists", exists))
	return exists, nil
}

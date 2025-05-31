package file_format_service

import (
	"context"
	"log/slog"
)

type FileFormatProvider interface {
	ListFileFormats(ctx context.Context) ([]string, error)
	FileFormatExists(ctx context.Context, format string) (bool, error)
}

type FileFormatService struct {
	provider FileFormatProvider
	logger   *slog.Logger
}

func New(provider FileFormatProvider, logger *slog.Logger) *FileFormatService {
	const op = "service.New.Library.FileFormats"

	if logger == nil {
		logger = slog.Default()
	}
	logger.Info("initializing library handle - service file formats", slog.String("op", op))

	return &FileFormatService{
		provider: provider,
		logger:   logger,
	}
}

package file_formats_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) FileFormatExists(ctx context.Context, format string) (bool, error) {
	const op = "repository.Library.FileFormats.FileFormatExists"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("checking format existence", slog.String("format", format))

	query := `SELECT EXISTS(SELECT 1 FROM file_formats WHERE format = $1)`

	var exists bool
	err := r.db.QueryRow(ctx, query, format).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("format not found", slog.String("format", format))
			return false, nil
		}
		logger.Error("database query failed", slog.String("error", err.Error()), slog.String("format", format))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("format existence checked", slog.String("format", format), slog.Bool("exists", exists))
	return exists, nil
}

package file_formats_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) ListFileFormats(ctx context.Context) ([]string, error) {
	const op = "repository.Library.FileFormats.ListFileFormats"
	logger := r.logger.With(slog.String("op", op))

	query := `SELECT format FROM file_formats ORDER BY format ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("no file formats found")
			return []string{}, nil
		}
		logger.Error("failed to query file formats", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	formats := make([]string, 0)
	for rows.Next() {
		var format string
		if err = rows.Scan(&format); err != nil {
			logger.Error("failed to scan file format", slog.String("error", err.Error()))
			return nil, fmt.Errorf("%s: scan error: %w", op, err)
		}
		formats = append(formats, format)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: rows error: %w", op, err)
	}

	logger.Debug("file formats listed", slog.Int("count", len(formats)))
	return formats, nil
}

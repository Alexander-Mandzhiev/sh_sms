package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) ExistByCode(ctx context.Context, code string, appID int) (bool, error) {
	const op = "repository.Permission.ExistByCode"
	logger := r.logger.With(slog.String("op", op), slog.String("code", code), slog.Int("app_id", appID))
	query := `SELECT EXISTS(SELECT 1 FROM permissions WHERE code = $1 AND app_id = $2 AND deleted_at IS NULL)`
	logger.Debug("executing existence check", slog.String("query", query))
	var exists bool
	err := r.db.QueryRow(ctx, query, code, appID).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("no matching records found")
			return false, nil
		}

		logger.Error("database query failed", slog.Any("error", err), slog.String("code", code))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("existence check result", slog.Bool("exists", exists), slog.String("code", code))
	return exists, nil
}

package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) ExistByID(ctx context.Context, id uuid.UUID, appID int) (bool, error) {
	const op = "repository.Permission.ExistByID"
	logger := r.logger.With(slog.String("op", op), slog.String("permission_id", id.String()), slog.Int("app_id", appID))
	query := `SELECT EXISTS(SELECT 1 FROM permissions WHERE id = $1 AND app_id = $2)`
	logger.Debug("executing existence check", slog.String("query", query), slog.Any("id", id), slog.Int("app_id", appID))
	var exists bool
	err := r.db.QueryRow(ctx, query, id, appID).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("no matching records found")
			return false, nil
		}

		logger.Error("database query failed", slog.Any("error", err), slog.String("id", id.String()))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("existence check result", slog.Bool("exists", exists), slog.String("id", id.String()))
	return exists, nil
}

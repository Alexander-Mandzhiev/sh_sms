package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) HardDelete(ctx context.Context, id uuid.UUID, appID int) error {
	const op = "repository.Permission.HardDelete"
	logger := r.logger.With(slog.String("op", op), slog.String("permission_id", id.String()), slog.Int("app_id", appID))
	query := `DELETE FROM permissions WHERE id = $1 AND app_id = $2`
	logger.Debug("executing hard delete", slog.String("query", query), slog.String("id", id.String()), slog.Int("app_id", appID))
	result, err := r.db.Exec(ctx, query, id, appID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
			return fmt.Errorf("%s: %w", op, ErrDatabase)
		}
		logger.Error("operation failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("no rows deleted - permission not found")
		return fmt.Errorf("%w: %s", pgx.ErrNoRows, op)
	}

	logger.Info("permission hard-deleted successfully", slog.Int64("rows_affected", rowsAffected))
	return nil
}

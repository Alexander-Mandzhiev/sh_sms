package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) SoftDelete(ctx context.Context, id uuid.UUID, appID int) error {
	const op = "repository.Permission.SoftDelete"
	logger := r.logger.With(slog.String("op", op), slog.String("permission_id", id.String()), slog.Int("app_id", appID))
	query := `UPDATE permissions SET deleted_at = NOW() WHERE id = $1 AND app_id = $2 AND deleted_at IS NULL`
	logger.Debug("executing soft delete", slog.String("query", query), slog.String("id", id.String()), slog.Int("app_id", appID))
	result, err := r.db.Exec(ctx, query, id, appID)
	if err != nil {
		logger.Error("database operation failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("no rows affected - permission not found or already deleted")
		return fmt.Errorf("%w: %s", pgx.ErrNoRows, op)
	}

	logger.Info("permission soft-deleted successfully", slog.Int64("rows_affected", rowsAffected))
	return nil
}

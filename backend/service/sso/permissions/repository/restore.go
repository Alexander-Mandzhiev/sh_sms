package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Restore(ctx context.Context, id uuid.UUID, appID int) error {
	const op = "repository.Permission.Restore"
	logger := r.logger.With(slog.String("op", op), slog.String("permission_id", id.String()), slog.Int("app_id", appID))
	query := `UPDATE permissions SET is_active = true, deleted_at = NULL WHERE id = $1 AND app_id = $2`
	result, err := r.db.Exec(ctx, query, id, appID)
	if err != nil {
		logger.Error("failed to execute restore query", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("no rows affected - permission not found or not deleted")
		return fmt.Errorf("%w: %s", pgx.ErrNoRows, op)
	}

	logger.Info("permission restored successfully", slog.Int64("rows_affected", rowsAffected))
	return nil
}

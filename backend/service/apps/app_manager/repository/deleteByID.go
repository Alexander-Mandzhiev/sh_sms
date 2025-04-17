package repository

import (
	"backend/service/apps/constants"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) DeleteByID(ctx context.Context, id int) error {
	const op = "repository.AppRepository.DeleteByID"
	logger := r.logger.With(slog.String("op", op), slog.Int("app_id", id))
	query := `UPDATE apps SET is_active = false, updated_at = CURRENT_TIMESTAMP WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Deactivation failed", slog.String("query", query), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("No active app found with ID")
		return constants.ErrNotFound
	}

	logger.Info("App deactivated successfully")
	return nil
}

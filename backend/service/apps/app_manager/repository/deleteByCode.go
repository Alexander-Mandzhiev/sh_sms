package repository

import (
	"backend/service/constants"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) DeleteByCode(ctx context.Context, code string) error {
	const op = "repository.AppRepository.DeleteByCode"
	logger := r.logger.With(slog.String("op", op), slog.String("app_code", code))
	query := `UPDATE apps SET is_active = false, updated_at = CURRENT_TIMESTAMP WHERE code = $1`
	result, err := r.db.Exec(ctx, query, code)
	if err != nil {
		logger.Error("Deactivation failed", slog.String("query", query), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("No active app found with code")
		return constants.ErrNotFound
	}

	logger.Info("App deactivated successfully")
	return nil
}

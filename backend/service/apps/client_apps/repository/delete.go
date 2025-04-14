package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, clientID string, appID int) error {
	const op = "repository.Delete"
	logger := r.logger.With(slog.String("op", op))
	query := `DELETE FROM client_apps WHERE client_id = $1 AND app_id = $2`

	result, err := r.db.Exec(ctx, query, clientID, appID)
	if err != nil {
		logger.Error("failed to delete client app", sl.Err(err, false))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("client app not found for deletion", slog.String("client_id", clientID), slog.Int("app_id", appID))
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	logger.Info("client app deleted", slog.String("client_id", clientID), slog.Int("app_id", appID))
	return nil
}

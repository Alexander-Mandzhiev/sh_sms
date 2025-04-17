package repository

import (
	"backend/service/apps/constants"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, clientID string, appID int) error {
	const op = "repository.Delete"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID))
	query := `UPDATE client_apps SET is_active = false, updated_at = NOW() WHERE client_id = $1 AND app_id = $2 RETURNING client_id`
	var deletedID string
	err := r.db.QueryRow(ctx, query, clientID, appID).Scan(&deletedID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("client app not found for deactivation")
			return fmt.Errorf("%s: %w", op, constants.ErrNotFound)
		}
		logger.Error("failed to deactivate client app", slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	logger.Info("client app deactivated successfully", slog.String("client_id", clientID), slog.Int("app_id", appID))
	return nil
}

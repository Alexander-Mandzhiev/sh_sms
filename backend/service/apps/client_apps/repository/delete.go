package repository

import (
	sl "backend/pkg/logger"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, clientID string, appID int32) error {
	const op = "repository.Delete"
	logger := r.logger.With(slog.String("op", op))
	query := `DELETE FROM client_apps WHERE client_id = $1 AND app_id = $2`

	result, err := r.db.Exec(ctx, query, clientID, appID)
	if err != nil {
		logger.Error("failed to delete client app", sl.Err(err, true))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

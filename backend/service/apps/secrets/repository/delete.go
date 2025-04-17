package repository

import (
	"backend/service/apps/constants"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, clientID string, appID int, secretType string) error {
	const op = "repository.Secret.Delete"
	logger := r.logger.With(slog.String("op", op))

	query := `DELETE FROM secrets WHERE client_id = $1 AND app_id = $2 AND secret_type = $3`
	result, err := r.db.Exec(ctx, query, clientID, appID, secretType)

	if err != nil {
		logger.Error("failed to delete secret", slog.Any("error", err), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("secret_type", secretType))
		return fmt.Errorf("delete secret failed: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("secret not found for deletion")
		return constants.ErrNotFound
	}

	logger.Debug("secret deleted successfully",
		slog.Int64("rows_affected", rowsAffected))

	return nil
}

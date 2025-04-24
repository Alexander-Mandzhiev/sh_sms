package repository

import (
	"backend/service/constants"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) HardDelete(ctx context.Context, clientID, roleID uuid.UUID) error {
	const op = "repository.Roles.HardDelete"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", roleID.String()))

	query := `DELETE FROM roles WHERE id = $1 AND client_id = $2`

	result, err := r.db.Exec(ctx, query, roleID, clientID)
	if err != nil {
		logger.Error("database error", slog.Any("error", err), slog.String("query", query))
		return fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	if rowsAffected := result.RowsAffected(); rowsAffected == 0 {
		logger.Warn("role not found for deletion")
		return fmt.Errorf("%w: role", constants.ErrNotFound)
	}

	logger.Info("role hard-deleted successfully")
	return nil
}

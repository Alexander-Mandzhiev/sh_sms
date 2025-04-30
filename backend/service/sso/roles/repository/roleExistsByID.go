package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) RoleExistsByID(ctx context.Context, clientID uuid.UUID, roleID uuid.UUID, appID int) (bool, error) {
	const op = "repository.Roles.RoleExistsByID"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", roleID.String()))
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE id = $1 AND client_id = $2 AND app_id = $3 AND deleted_at IS NULL)`
	var exists bool
	err := r.db.QueryRow(ctx, query, roleID, clientID, appID).Scan(&exists)
	if err != nil {
		logger.Error("database error", slog.Any("error", err), slog.Int("app_id", appID))
		return false, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("role existence checked", slog.Bool("exists", exists))
	return exists, nil
}

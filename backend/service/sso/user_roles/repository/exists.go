package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Exists(ctx context.Context, userID uuid.UUID, roleID uuid.UUID, clientID uuid.UUID, appID int) (bool, error) {
	const op = "repository.UserRoles.Exists"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()),
		slog.String("role_id", roleID.String()), slog.String("client_id", clientID.String()), slog.Int("app_id", appID))
	query := `SELECT EXISTS(SELECT 1 FROM user_roles WHERE user_id = $1 AND role_id = $2 AND client_id = $3 AND app_id = $4)`

	var exists bool
	err := r.db.QueryRow(ctx, query, userID, roleID, clientID, appID).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("relation user-role not found")
			return false, nil
		}
		logger.Error("database error", slog.Any("error", err))
		return false, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("existence check completed", slog.Bool("exists", exists))
	return exists, nil
}

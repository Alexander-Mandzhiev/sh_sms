package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) RoleExists(ctx context.Context, clientID uuid.UUID, appID int, name string) (bool, error) {
	const op = "repository.Roles.RoleExists"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("name", name))
	query := `SELECT 1 FROM roles WHERE client_id = $1 AND app_id = $2 AND name = $3 AND deleted_at IS NULL LIMIT 1`

	var exists int
	err := r.db.QueryRow(ctx, query, clientID, appID, name).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("role not found")
			return false, nil
		}
		logger.Error("database error", slog.String("query", query), slog.Any("error", err))
		return false, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("role exists")
	return true, nil
}

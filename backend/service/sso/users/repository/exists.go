package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) Exists(ctx context.Context, clientID, userID uuid.UUID) (bool, error) {
	const op = "repository.User.Exists"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("checking user existence")

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND client_id = $2)`

	var exists bool
	if err := r.db.QueryRow(ctx, query, userID, clientID).Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Debug("user not found")
			return false, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("database error", slog.Any("error", err))
		return false, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("existence check result", slog.Bool("exists", exists))
	return exists, nil
}

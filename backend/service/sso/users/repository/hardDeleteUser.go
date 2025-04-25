package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) HardDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error {
	const op = "service.User.HardDeleteUser"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to hard delete user")
	query := `DELETE FROM users WHERE id = $1 AND client_id = $2 RETURNING id`

	var deletedID uuid.UUID
	err := r.db.QueryRow(ctx, query, userID, clientID).Scan(&deletedID)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Warn("user not found for deletion")
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("database operation failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("user permanently deleted", slog.String("deleted_id", deletedID.String()))
	return nil
}

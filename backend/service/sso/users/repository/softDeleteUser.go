package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) SoftDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error {
	const op = "repository.User.SoftDeleteUser"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to soft delete user")
	query := `UPDATE users SET is_active = false, deleted_at = NOW(), updated_at = NOW() WHERE id = $1 AND client_id = $2 RETURNING id`
	var deletedID uuid.UUID
	err := r.db.QueryRow(ctx, query, userID, clientID).Scan(&deletedID)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Warn("user not found or already deleted")
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("database operation failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("user soft-deleted successfully")
	return nil
}

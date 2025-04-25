package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) UpdatePasswordHash(ctx context.Context, userID uuid.UUID, passwordHash string) error {
	const op = "repository.User.UpdatePasswordHash"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()))
	logger.Debug("attempting to update password hash")

	query := `UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2 RETURNING id`
	result, err := r.db.Exec(ctx, query, passwordHash, userID)
	if err != nil {
		logger.Error("database operation failed", slog.Any("error", err), slog.String("query", query))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("no rows affected - user not found")
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	logger.Info("password hash updated successfully", slog.Int64("rows_affected", rowsAffected))
	return nil
}

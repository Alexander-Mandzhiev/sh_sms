package repository

import (
	"backend/service/sso/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

func (r *Repository) Update(ctx context.Context, user *models.User) error {
	const op = "repository.User.Update"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", user.ID.String()), slog.String("client_id", user.ClientID.String()))
	query := `UPDATE users SET email = $1, full_name = $2, phone = $3, updated_at = NOW() WHERE id = $4 AND client_id = $5 RETURNING updated_at`
	logger.Debug("executing update query", slog.String("query", query))

	var updatedAt time.Time
	if err := r.db.QueryRow(ctx, query, user.Email, user.FullName, user.Phone, user.ID, user.ClientID).Scan(&updatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warn("user not found or no changes")
			return fmt.Errorf("%w: %v", ErrNotFound, err)
		}
		logger.Error("database operation failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	user.UpdatedAt = updatedAt
	logger.Info("user updated successfully", slog.Time("new_updated_at", updatedAt))
	return nil
}

package repository

import (
	"backend/service/sso/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Restore(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error) {
	const op = "repository.User.Restore"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to restore user")

	query := `UPDATE users SET deleted_at = NULL, is_active = TRUE, updated_at = NOW() WHERE id = $1 AND client_id = $2
        RETURNING id, client_id, email, full_name, phone, is_active, created_at, updated_at, deleted_at`

	logger.Debug("executing restore query", query)

	var user models.User
	var deletedAt sql.NullTime

	if err := r.db.QueryRow(ctx, query, userID, clientID).
		Scan(&user.ID, &user.ClientID, &user.Email, &user.FullName, &user.Phone, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &deletedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("user not found or already restored")
			return nil, fmt.Errorf("%w", ErrNotFound)
		}
		logger.Error("database operation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if deletedAt.Valid {
		user.DeletedAt = &deletedAt.Time
	} else {
		user.DeletedAt = nil
	}

	logger.Info("user restored successfully", slog.Time("restored_at", user.UpdatedAt))
	return &user, nil
}

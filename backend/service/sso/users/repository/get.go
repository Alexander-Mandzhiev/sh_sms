package repository

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error) {
	const op = "repository.User.Get"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to get user")
	query := `SELECT id, client_id, email, password_hash, full_name, phone, is_active, created_at, updated_at FROM users WHERE id = $1 AND client_id = $2`

	var user models.User
	err := r.db.QueryRow(ctx, query, userID, clientID).Scan(&user.ID, &user.ClientID, &user.Email, &user.PasswordHash, &user.FullName, &user.Phone, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("user not found")
			return nil, fmt.Errorf("%s: %w", op, constants.ErrNotFound)
		}
		logger.Error("database query failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("user found")
	return &user, nil
}

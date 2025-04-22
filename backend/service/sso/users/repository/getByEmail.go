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

func (r *Repository) GetByEmail(ctx context.Context, clientID uuid.UUID, email string) (*models.User, error) {
	const op = "repository.User.GetByEmail"
	logger := r.logger.With(slog.String("op", op), slog.String("email", email), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to get user by email")
	query := `SELECT id, client_id, email, password_hash, full_name, phone, is_active, created_at, updated_at FROM users WHERE email = $1 AND client_id = $2 AND deleted_at IS NULL`

	var user models.User
	err := r.db.QueryRow(ctx, query, email, clientID).Scan(&user.ID, &user.ClientID, &user.Email, &user.PasswordHash, &user.FullName, &user.Phone, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("user not found by email")
			return nil, fmt.Errorf("%s: %w", op, constants.ErrNotFound)
		}
		logger.Error("database query failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("user found by email")
	return &user, nil
}

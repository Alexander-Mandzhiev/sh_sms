package repository

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetByEmailOrPhone(ctx context.Context, clientID uuid.UUID, login string) (*models.User, error) {
	const op = "repository.User.GetByLogin"
	logger := r.logger.With(slog.String("op", op), slog.String("login", login), slog.String("client_id", clientID.String()))

	logger.Debug("attempting to get user by login (email or phone)")

	query := `SELECT id, client_id, email, password_hash, full_name, phone, is_active, created_at, updated_at, deleted_at 
        FROM users WHERE (email = $1 OR phone = $1) AND client_id = $2 LIMIT 1`

	var user models.User
	err := r.db.QueryRow(ctx, query, login, clientID).
		Scan(&user.ID, &user.ClientID, &user.Email, &user.PasswordHash, &user.FullName, &user.Phone, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("user not found by login")
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("database query failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("user found by login", slog.String("user_id", user.ID.String()))
	return &user, nil
}

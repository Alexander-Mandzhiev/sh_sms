package repository

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, user *models.User) error {
	const op = "repository.User.Create"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", user.ID.String()), slog.String("client_id", user.ClientID.String()), slog.String("email", user.Email))
	logger.Debug("attempting to create user")

	if user.Email == "" || user.PasswordHash == "" {
		return fmt.Errorf("missing required fields")
	}

	query := `INSERT INTO users (id, client_id, email, password_hash, full_name, phone, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.Exec(ctx, query, user.ID, user.ClientID, user.Email, user.PasswordHash, user.FullName, user.Phone, user.IsActive, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		logger.Error("failed to create user", slog.Any("error", err))
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "idx_users_email_client" {
				return fmt.Errorf("%s: %w", op, constants.ErrEmailAlreadyExists)
			}
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("user created successfully")
	return nil
}

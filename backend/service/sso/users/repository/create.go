package repository

import (
	"backend/service/sso/models"
	"context"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, user *models.User) error {
	const op = "repository.User.Create"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", user.ID.String()))
	logger.Debug("attempting to create user")
	return nil
}

package repository

import (
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) Get(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error) {
	const op = "repository.User.Get"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to get user")
	return nil, nil
}

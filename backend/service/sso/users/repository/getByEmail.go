package repository

import (
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) GetByEmail(ctx context.Context, clientID uuid.UUID, email string) (*models.User, error) {
	const op = "repository.User.GetByEmail"
	logger := r.logger.With(slog.String("op", op), slog.String("email", email), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to get user")
	return nil, nil
}

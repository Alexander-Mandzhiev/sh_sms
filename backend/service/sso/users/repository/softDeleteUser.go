package repository

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) SoftDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error {
	const op = "service.User.SoftDeleteUser"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to soft delete user")
	return nil
}

package repository

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) HardDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error {
	const op = "service.User.HardDeleteUser"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to hard delete user")
	return nil
}

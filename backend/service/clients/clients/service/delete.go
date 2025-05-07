package service

import (
	"backend/service/clients/clients/handle"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID, permanent bool) error {
	const op = "service.Clients.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("id", id.String()), slog.Bool("permanent", permanent))
	logger.Debug("processing client deleted")

	if err := utils.ValidateUUID(id); err != nil {
		logger.Warn("invalid client id provided")
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := s.provider.Delete(ctx, id, permanent); err != nil {
		if errors.Is(err, handle.ErrNotFound) {
			logger.Warn("client not found")
			return handle.ErrNotFound
		}
		logger.Error("database error", slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, handle.ErrInternal)
	}

	logger.Info("client deleted", slog.String("id", id.String()), slog.Bool("permanent", permanent))
	return nil
}

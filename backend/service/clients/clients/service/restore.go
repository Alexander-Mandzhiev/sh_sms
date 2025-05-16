package service

import (
	"backend/pkg/utils"
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Restore(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	const op = "service.Clients.Restore"
	logger := s.logger.With(slog.String("op", op), slog.String("id", id.String()))
	logger.Debug("processing client restored")

	if err := utils.ValidateUUID(id); err != nil {
		logger.Warn("invalid client id provided")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	client, err := s.provider.Restore(ctx, id)
	if err != nil {
		if errors.Is(err, handle.ErrNotFound) {
			logger.Warn("client not found")
			return nil, handle.ErrNotFound
		}
		if errors.Is(err, handle.ErrCodeExists) {
			logger.Warn("duplicate entry during restore")
			return nil, err
		}
		if errors.Is(err, handle.ErrInvalidArgument) {
			logger.Warn("invalid type_id during restore")
			return nil, err
		}
		logger.Error("database error", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, handle.ErrInternal)
	}

	logger.Debug("client restored successfully")
	return client, nil
}

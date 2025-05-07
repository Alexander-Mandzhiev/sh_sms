package service

import (
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	const op = "service.Clients.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("id", id.String()))
	logger.Debug("processing client creation")

	if err := utils.ValidateUUID(id); err != nil {
		logger.Warn("invalid client id provided")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	client, err := s.provider.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, handle.ErrNotFound) {
			logger.Warn("client not found")
			return nil, handle.ErrNotFound
		}
		logger.Error("database error", slog.String("error", err.Error()))
		return nil, handle.ErrInternal
	}

	logger.Debug("client retrieved successfully")
	return client, nil
}

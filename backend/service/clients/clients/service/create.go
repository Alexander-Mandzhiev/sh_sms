package service

import (
	"backend/pkg/utils"
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, param *models.CreateParams) (*models.Client, error) {
	const op = "service.Clients.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("name", param.Name), slog.Int("type_id", param.TypeID))
	logger.Debug("processing client creation")

	if err := utils.ValidateString(param.Name, 255); err != nil {
		logger.Warn("invalid name", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%w: %s", err, op)
	}

	if err := utils.ValidateString(param.Description, 10000); err != nil {
		logger.Warn("invalid description", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%w: %s", err, op)
	}

	if err := utils.ValidateWebsite(param.Website, 255); err != nil {
		logger.Warn("invalid website", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%w: %s", err, op)
	}

	if param.TypeID <= 0 {
		logger.Warn("invalid type_id", slog.Int("value", param.TypeID))
		return nil, fmt.Errorf("%w: type_id must be positive", handle.ErrInvalidArgument)
	}

	client, err := s.provider.Create(ctx, param)
	if err != nil {
		logger.Warn("failed to create client", slog.String("error", err.Error()))

		if errors.Is(err, handle.ErrCodeExists) || errors.Is(err, handle.ErrInvalidArgument) {
			return nil, err
		}
		return nil, handle.ErrInternal
	}

	logger.Info("client created")
	return client, nil
}

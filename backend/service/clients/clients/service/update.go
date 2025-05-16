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

func (s *Service) Update(ctx context.Context, params *models.UpdateParams) (*models.Client, error) {
	const op = "service.Clients.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ID))
	logger.Debug("processing client update")

	if params.Name == nil && params.Description == nil && params.TypeID == nil && params.Website == nil {
		logger.Warn("no fields to update")
		return nil, handle.ErrInvalidArgument
	}

	if err := utils.ValidateUUIDToString(params.ID); err != nil {
		logger.Warn("invalid client ID format")
		return nil, handle.ErrInvalidArgument
	}

	if params.Name != nil && *params.Name == "" {
		params.Name = nil
	}
	if params.Description != nil && *params.Description == "" {
		params.Description = nil
	}
	if params.Website != nil && *params.Website == "" {
		params.Website = nil
	}

	client, err := s.provider.Update(ctx, params)
	if err != nil {
		if errors.Is(err, handle.ErrNotFound) {
			logger.Warn("client not found")
			return nil, handle.ErrNotFound
		}
		if errors.Is(err, handle.ErrCodeExists) {
			logger.Warn("duplicate client name")
			return nil, handle.ErrCodeExists
		}
		if errors.Is(err, handle.ErrInvalidArgument) {
			logger.Warn("invalid type_id")
			return nil, handle.ErrInvalidArgument
		}
		logger.Error("database error", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, handle.ErrInternal)

	}

	logger.Debug("client updated successfully")
	return client, nil
}

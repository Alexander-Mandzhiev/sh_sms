package service

import (
	"backend/service/apps/client_apps/handle"
	"backend/service/apps/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, params models.CreateClientApp) (*models.ClientApp, error) {
	const op = "service.ClientApp.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID), slog.Int("app_id", params.AppID))

	if err := validateClientID(params.ClientID); err != nil {
		logger.Warn("client ID validation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
	}

	if params.AppID <= 0 {
		err := fmt.Errorf("invalid app_id: %d", params.AppID)
		logger.Warn("app ID validation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
	}

	clientApp, err := s.provider.Create(ctx, params)
	if err != nil {
		logger.Error("failed to create client app", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client app created", slog.Bool("is_active", clientApp.IsActive), slog.Time("created_at", clientApp.CreatedAt))
	return clientApp, nil
}

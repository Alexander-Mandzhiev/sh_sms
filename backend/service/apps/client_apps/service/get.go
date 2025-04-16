package service

import (
	sl "backend/pkg/logger"
	"backend/service/apps/client_apps/handle"
	"backend/service/apps/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, clientID string, appID int) (*models.ClientApp, error) {
	const op = "service.ClientApp.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID))

	if err := validateClientID(clientID); err != nil {
		logger.Warn("client ID validation failed", sl.Err(err, false))
		return nil, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
	}

	if appID <= 0 {
		err := fmt.Errorf("invalid app_id: %d", appID)
		logger.Warn("app ID validation failed", sl.Err(err, false))
		return nil, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
	}

	clientApp, err := s.provider.Get(ctx, clientID, appID)
	if err != nil {
		logger.Error("failed to get client app", sl.Err(err, false))
		return nil, s.convertError(err)
	}

	logger.Info("client app retrieved successfully", slog.Bool("is_active", clientApp.IsActive), slog.Time("updated_at", clientApp.UpdatedAt))

	return clientApp, nil
}

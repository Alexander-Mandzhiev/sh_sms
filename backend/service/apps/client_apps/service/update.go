package service

import (
	"backend/service/apps/client_apps/handle"
	"backend/service/apps/models"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (s *Service) Update(ctx context.Context, clientID string, appID int, isActive *bool) (*models.ClientApp, error) {
	const op = "service.ClientApp.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID))

	if err := validateClientID(clientID); err != nil {
		logger.Warn("client ID validation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
	}

	if appID <= 0 {
		err := fmt.Errorf("invalid app_id: %d", appID)
		logger.Warn("app ID validation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
	}

	updateParams := models.UpdateClientApp{
		ClientID:  clientID,
		AppID:     appID,
		IsActive:  isActive,
		UpdatedAt: time.Now().UTC(),
	}

	updatedApp, err := s.provider.Update(ctx, updateParams)
	if err != nil {
		logger.Error("failed to update client app", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client app updated successfully", slog.Bool("new_is_active", updatedApp.IsActive), slog.Time("updated_at", updatedApp.UpdatedAt))
	return updatedApp, nil
}

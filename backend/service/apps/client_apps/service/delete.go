package service

import (
	sl "backend/pkg/logger"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, clientID string, appID int) error {
	const op = "service.ClientApp.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID))

	if err := utils.ValidateUUIDToString(clientID); err != nil {
		logger.Warn("client ID validation failed", sl.Err(err, false))
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		err = fmt.Errorf("invalid app_id: %d", appID)
		logger.Warn("app ID validation failed", sl.Err(err, false))
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if err := s.provider.Delete(ctx, clientID, appID); err != nil {
		logger.Error("failed to delete client app", sl.Err(err, false))
		return s.convertError(err)
	}

	logger.Info("client app deleted successfully")
	return nil
}

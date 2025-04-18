package service

import (
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, clientID string, appID int, secretType string) error {
	const op = "service.Secret.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("secret_type", secretType))
	logger.Debug("starting secret deletion")

	if err := utils.ValidateClientID(clientID); err != nil {
		logger.Warn("invalid client ID", slog.Any("error", err))
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if !utils.IsValidSecretType(secretType) {
		logger.Warn("invalid secret type")
		return fmt.Errorf("%w: invalid secret type", constants.ErrInvalidArgument)
	}

	if err := s.provider.Delete(ctx, clientID, appID, secretType); err != nil {
		logger.Error("delete transaction failed", slog.Any("error", err))
		return s.convertError(err)
	}

	logger.Info("secret deleted successfully")
	return nil
}

package service

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error) {
	const op = "service.Secret.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("secret_type", secretType))
	logger.Debug("starting secret retrieval")

	if err := utils.ValidateUUIDToString(clientID); err != nil {
		logger.Warn("invalid client ID", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if !utils.IsValidSecretType(secretType) {
		logger.Warn("invalid secret type")
		return nil, fmt.Errorf("%w: invalid secret type", constants.ErrInvalidArgument)
	}

	secret, err := s.provider.Get(ctx, clientID, appID, secretType)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("secret not found")
			return nil, fmt.Errorf("%w: %v", constants.ErrNotFound, err)
		}
		logger.Error("failed to get secret", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	if secret.RevokedAt != nil {
		logger.Info("retrieved revoked secret", slog.Time("revoked_at", *secret.RevokedAt))
	} else {
		logger.Debug("retrieved active secret")
	}

	logger.Info("secret retrieved successfully", slog.Int("version", secret.SecretVersion), slog.Time("generated_at", secret.GeneratedAt))
	return secret, nil
}

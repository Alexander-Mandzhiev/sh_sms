package service

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Rotate(ctx context.Context, params models.RotateSecretParams) (*models.Secret, error) {
	const op = "service.Secret.Rotate"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID), slog.Int("app_id", params.AppID), slog.String("secret_type", params.SecretType), slog.String("rotated_by", params.RotatedBy))
	logger.Debug("starting secret rotation")

	if err := utils.ValidateClientID(params.ClientID); err != nil {
		logger.Warn("invalid client ID", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if err := utils.ValidateAppID(params.AppID); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if !utils.IsValidSecretType(params.SecretType) {
		logger.Warn("invalid secret type")
		return nil, fmt.Errorf("%w: invalid secret type", constants.ErrInvalidArgument)
	}

	if params.RotatedBy != "" {
		if err := utils.ValidateClientID(params.RotatedBy); err != nil {
			logger.Warn("invalid rotated_by", slog.Any("error", err))
			return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
	}

	secret, err := s.provider.Rotate(ctx, params)
	if err != nil {
		logger.Error("rotation transaction failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("secret rotated successfully", slog.Int("new_version", secret.SecretVersion), slog.Bool("is_active", secret.RevokedAt == nil))

	return secret, nil
}

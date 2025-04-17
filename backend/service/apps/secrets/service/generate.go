package service

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Generate(ctx context.Context, params models.CreateSecretParams) (*models.Secret, error) {
	const op = "service.Secret.Generate"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID), slog.Int("app_id", params.AppID), slog.String("secret_type", params.SecretType))
	logger.Debug("starting secret generation")

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

	if params.Algorithm == "" {
		params.Algorithm = "bcrypt"
		logger.Debug("using default algorithm")
	}

	secret, err := s.provider.Create(ctx, params)
	if err != nil {
		if errors.Is(err, constants.ErrSecretAlreadyExists) {
			logger.Warn("secret already exists")
			return nil, fmt.Errorf("%w: %v", constants.ErrSecretAlreadyExists, err)
		}
		logger.Error("failed to generate secret", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("secret generated successfully")
	return secret, nil
}

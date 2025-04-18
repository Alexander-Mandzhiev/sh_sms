package service

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Revoke(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error) {
	const op = "service.Secret.Revoke"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("secret_type", secretType))
	logger.Debug("starting secret revocation")

	if err := utils.ValidateClientID(clientID); err != nil {
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

	secret, err := s.provider.Revoke(ctx, clientID, appID, secretType)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "secret not found")
		}
		if errors.Is(err, constants.ErrAlreadyRevoked) {
			return nil, status.Error(codes.FailedPrecondition, "secret already revoked")
		}
		return nil, s.convertError(err)
	}

	logger.Info("secret revoked successfully", slog.Time("revoked_at", *secret.RevokedAt), slog.Int("version", secret.SecretVersion))
	return secret, nil
}

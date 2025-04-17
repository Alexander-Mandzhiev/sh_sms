package service

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

func (s *Service) GetRotation(ctx context.Context, clientID string, appID int, secretType string, rotatedAt time.Time) (*models.RotationHistory, error) {
	const op = "service.Secret.GetRotation"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID),
		slog.String("secret_type", secretType), slog.Time("rotated_at", rotatedAt),
	)
	logger.Debug("starting rotation history retrieval")

	if err := utils.ValidateClientID(clientID); err != nil {
		logger.Warn("invalid client ID", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if appID <= 0 {
		logger.Warn("invalid app ID")
		return nil, fmt.Errorf("%w: app_id must be positive", constants.ErrInvalidArgument)
	}

	if !utils.IsValidSecretType(secretType) {
		logger.Warn("invalid secret type")
		return nil, fmt.Errorf("%w: invalid secret type", constants.ErrInvalidArgument)
	}

	if rotatedAt.IsZero() {
		logger.Warn("zero time for rotated_at")
		return nil, fmt.Errorf("%w: rotated_at is required", constants.ErrInvalidArgument)
	}

	if rotatedAt.After(time.Now()) {
		logger.Warn("future rotated_at time")
		return nil, fmt.Errorf("%w: rotated_at cannot be in future", constants.ErrInvalidArgument)
	}

	normalizedSecretType := strings.ToLower(secretType)
	if normalizedSecretType != secretType {
		logger.Debug("normalized secret type", slog.String("original", secretType), slog.String("normalized", normalizedSecretType))
		secretType = normalizedSecretType
	}

	history, err := s.provider.GetRotation(ctx, clientID, appID, secretType, rotatedAt)
	if err != nil {
		logger.Error("failed to get rotation history", slog.Any("error", err), slog.Time("rotated_at", rotatedAt))
		return nil, s.convertError(err)
	}

	if history == nil {
		logger.Error("unexpected nil history")
		return nil, s.convertError(constants.ErrInternal)
	}

	logger.Info("rotation history retrieved successfully", slog.String("old_secret_prefix", maskSecret(history.OldSecret)),
		slog.String("new_secret_prefix", maskSecret(history.NewSecret)))
	return history, nil
}

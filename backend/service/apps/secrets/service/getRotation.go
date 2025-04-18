package service

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) GetRotation(ctx context.Context, id int) (*models.RotationHistory, error) {
	const op = "service.Secret.GetRotation"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", id))
	logger.Debug("starting rotation history retrieval")

	if id <= 0 {
		logger.Warn("invalid rotation ID")
		return nil, fmt.Errorf("%w: invalid rotation ID", constants.ErrInvalidArgument)
	}

	history, err := s.provider.GetRotation(ctx, id)
	if err != nil {
		logger.Error("failed to get rotation history", slog.Any("error", err), slog.Int("id", id))
		return nil, s.convertError(err)
	}

	if history == nil {
		logger.Error("empty rotation history response")
		return nil, s.convertError(constants.ErrInternal)
	}

	if err = utils.ValidateRotationHistory(history); err != nil {
		logger.Error("invalid rotation history data", slog.Any("error", err), slog.Any("history", maskRotationData(history)))
		return nil, s.convertError(constants.ErrInternal)
	}

	logger.Info("rotation history retrieved successfully", slog.String("client_id", history.ClientID), slog.Int("app_id", history.AppID), slog.String("secret_type", history.SecretType), slog.Time("rotated_at", history.RotatedAt))
	return history, nil
}

func maskRotationData(h *models.RotationHistory) slog.Value {
	return slog.GroupValue(
		slog.String("client_id", h.ClientID),
		slog.Int("app_id", h.AppID),
		slog.String("secret_type", h.SecretType),
		slog.String("old_secret", maskSecret(h.OldSecret)),
		slog.String("new_secret", maskSecret(h.NewSecret)),
		slog.Time("rotated_at", h.RotatedAt))
}

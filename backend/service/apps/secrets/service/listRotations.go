package service

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

func (s *Service) ListRotations(ctx context.Context, filter models.ListFilter) ([]*models.RotationHistory, int, error) {
	const op = "service.Secret.ListRotations"
	logger := s.logger.With(slog.String("op", op), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	logger.Debug("starting rotation history listing")

	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid pagination parameters", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if filter.ClientID != nil {
		if err := utils.ValidateClientID(*filter.ClientID); err != nil {
			logger.Warn("invalid client ID in filter", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
	}

	if filter.AppID != nil {
		if err := utils.ValidateAppID(*filter.AppID); err != nil {
			logger.Warn("invalid app ID in filter", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
	}

	if filter.SecretType != nil {
		normalized := strings.ToLower(*filter.SecretType)
		if normalized != *filter.SecretType {
			logger.Debug("normalizing secret type", slog.String("original", *filter.SecretType), slog.String("normalized", normalized))
			filter.SecretType = &normalized
		}

		if !utils.IsValidSecretType(*filter.SecretType) {
			logger.Warn("invalid secret type in filter")
			return nil, 0, fmt.Errorf("%w: invalid secret type", constants.ErrInvalidArgument)
		}
	}

	if filter.RotatedBy != nil {
		if err := utils.ValidateClientID(*filter.RotatedBy); err != nil {
			logger.Warn("invalid rotated_by in filter", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
	}

	if filter.RotatedAfter != nil && filter.RotatedBefore != nil {
		if filter.RotatedAfter.After(*filter.RotatedBefore) {
			logger.Warn("invalid time range: rotated_after > rotated_before")
			return nil, 0, fmt.Errorf("%w: invalid time range", constants.ErrInvalidArgument)
		}
	}

	if filter.RotatedAfter != nil && filter.RotatedAfter.After(time.Now()) {
		logger.Warn("rotated_after in future")
		return nil, 0, fmt.Errorf("%w: rotated_after cannot be in future", constants.ErrInvalidArgument)
	}

	if filter.RotatedBefore != nil && filter.RotatedBefore.After(time.Now()) {
		logger.Warn("rotated_before in future")
		return nil, 0, fmt.Errorf("%w: rotated_before cannot be in future", constants.ErrInvalidArgument)
	}

	history, total, err := s.provider.ListRotations(ctx, filter)
	if err != nil {
		logger.Error("failed to list rotation history", slog.Any("error", err))
		return nil, 0, s.convertError(err)
	}

	logger.Info("rotation history listed successfully", slog.Int("total_count", total), slog.Int("returned_count", len(history)))
	return history, total, nil
}

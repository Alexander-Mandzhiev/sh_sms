package service

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) List(ctx context.Context, filter models.ListFilter) ([]*models.Secret, int, error) {
	const op = "service.Secret.List"
	logger := s.logger.With(slog.String("op", op), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	logger.Debug("starting secrets listing")

	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid pagination parameters")
		return nil, 0, fmt.Errorf("%w: page and count must be positive", constants.ErrInvalidArgument)
	}

	if filter.ClientID != nil {
		if err := utils.ValidateClientID(*filter.ClientID); err != nil {
			logger.Warn("invalid client ID in filter", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
	}

	if filter.SecretType != nil && !utils.IsValidSecretType(*filter.SecretType) {
		logger.Warn("invalid secret type in filter")
		return nil, 0, fmt.Errorf("%w: invalid secret type", constants.ErrInvalidArgument)
	}

	if filter.RotatedBy != nil {
		if err := utils.ValidateClientID(*filter.RotatedBy); err != nil {
			logger.Warn("invalid rotated_by in filter", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
	}

	secrets, total, err := s.provider.List(ctx, filter)
	if err != nil {
		logger.Error("failed to list secrets", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	logger.Info("secrets list retrieved successfully", slog.Int("total_count", total), slog.Int("returned_count", len(secrets)))
	return secrets, total, nil
}

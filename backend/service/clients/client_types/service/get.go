package service

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, id int) (*models.ClientType, error) {
	const op = "service.ClientType.Get"
	logger := s.logger.With(slog.String("op", op), slog.Int("requested_id", id))
	logger.Debug("processing get request")

	if id <= 0 {
		err := fmt.Errorf("%w: invalid ID format", ErrInvalidArgument)
		logger.Warn("validation failed", slog.Any("error", err))
		return nil, err
	}

	ct, err := s.provider.Get(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("client type not found", slog.Int("id", id))
			return nil, fmt.Errorf("%w: %v", ErrNotFound, id)
		}
		logger.Error("database operation failed", slog.Any("error", err), slog.Int("id", id))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	if ct == nil {
		logger.Error("unexpected nil client type", slog.Int("id", id))
		return nil, fmt.Errorf("%w: empty response", ErrInternal)
	}

	logger.Info("client type retrieved successfully", slog.Int("id", ct.ID), slog.String("code", ct.Code))
	return ct, nil
}

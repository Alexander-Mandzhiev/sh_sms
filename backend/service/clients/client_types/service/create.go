package service

import (
	"backend/service/clients/client_types/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, ct *models.CreateParams) (*models.ClientType, error) {
	const op = "service.ClientType.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("code", ct.Code))
	logger.Debug("attempting to create client type")

	if err := utils.ValidateString(ct.Code, 50); err != nil {
		logger.Warn("code validation failed", slog.Any("error", err), slog.Int("code_length", len(ct.Code)))
		return nil, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	if err := utils.ValidateString(ct.Name, 100); err != nil {
		logger.Warn("name validation failed", slog.Any("error", err), slog.Int("name_length", len(ct.Name)))
		return nil, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	exists, err := s.provider.CodeExists(ctx, ct.Code)
	if err != nil {
		logger.Error("code existence check failed", slog.Any("error", err), slog.String("code", ct.Code))
		return nil, fmt.Errorf("%w: failed to check code uniqueness", ErrInternal)
	}
	if exists {
		logger.Warn("duplicate code", slog.String("code", ct.Code))
		return nil, fmt.Errorf("%w: %v", ErrCodeExists, ct.Code)
	}

	if ct.IsActive == nil {
		defaultActive := true
		ct.IsActive = &defaultActive
	}

	createdCt, err := s.provider.Create(ctx, ct)
	if err != nil {
		logger.Error("database operation failed", slog.Any("error", err), slog.String("code", ct.Code))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("client type created successfully", slog.Int("id", createdCt.ID),
		slog.Bool("is_active", createdCt.IsActive), slog.String("code", createdCt.Code))
	return createdCt, nil
}

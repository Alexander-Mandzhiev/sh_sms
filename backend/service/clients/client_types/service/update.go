package service

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, params *models.UpdateParams) (*models.ClientType, error) {
	const op = "service.ClientType.Update"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(params.ID)))
	logger.Debug("processing update", slog.String("new_code", params.Code), slog.String("new_name", params.Name), slog.String("new_description", params.Description))

	if params.ID <= 0 {
		logger.Warn("validation failed: invalid ID", slog.Int("received_id", int(params.ID)))
		return nil, fmt.Errorf("%w: invalid ID", ErrInvalidArgument)
	}

	if params.Code == "" || len(params.Code) > 50 {
		logger.Warn("validation failed: invalid code", slog.String("code", params.Code), slog.Int("code_length", len(params.Code)))
		return nil, fmt.Errorf("%w: code must be 1-50 chars", ErrInvalidArgument)
	}

	if params.Name == "" || len(params.Name) > 100 {
		logger.Warn("validation failed: invalid name", slog.String("name", params.Name), slog.Int("name_length", len(params.Name)))
		return nil, fmt.Errorf("%w: name must be 1-100 chars", ErrInvalidArgument)
	}

	existing, err := s.provider.Exist(ctx, int(params.ID))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("client type not found", slog.Int("id", int(params.ID)))
			return nil, fmt.Errorf("%w: %d", ErrNotFound, params.ID)
		}
		logger.Error("database check failed", slog.Any("error", err), slog.String("error_type", "database_error"))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	if !existing {
		logger.Warn("client type does not exist", slog.Int("id", int(params.ID)))
		return nil, fmt.Errorf("%w: %d", ErrNotFound, params.ID)
	}

	codeExists, err := s.provider.CodeExists(ctx, params.Code)
	if err != nil {
		logger.Error("code existence check failed", slog.String("code", params.Code), slog.Any("error", err))
		return nil, fmt.Errorf("%w: code check failed", ErrInternal)
	}

	if codeExists {
		logger.Warn("code conflict detected", slog.String("code", params.Code), slog.Int("id", int(params.ID)))
		return nil, fmt.Errorf("%w: code %s", ErrCodeConflict, params.Code)
	}

	updated, err := s.provider.Update(ctx, params)
	if err != nil {
		logger.Error("update operation failed", slog.Any("error", err), slog.Any("params", params))
		return nil, fmt.Errorf("%w: update failed", ErrInternal)
	}

	logger.Info("client type successfully updated", slog.String("new_code", updated.Code), slog.String("new_name", updated.Name), slog.String("new_description", updated.Description), slog.Bool("is_active", updated.IsActive))
	return updated, nil
}

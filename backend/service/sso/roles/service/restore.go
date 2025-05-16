package service

import (
	"backend/pkg/utils"
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Restore(ctx context.Context, clientID, roleID uuid.UUID, appID int) (*models.Role, error) {
	const op = "service.Roles.Restore"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", roleID.String()))
	logger.Debug("attempting to restore role")

	if err := utils.ValidateUUID(clientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(roleID); err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	role, err := s.provider.Restore(ctx, clientID, roleID, appID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("role not found")
			return nil, fmt.Errorf("%w: %s", ErrNotFound, "role")
		}
		logger.Error("database error", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Debug("successfully retrieved role", slog.String("role_name", role.Name))
	return role, nil
}

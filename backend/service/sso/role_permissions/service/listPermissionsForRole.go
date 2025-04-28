package service

import (
	serviceRole "backend/service/sso/roles/service"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) ListPermissionsForRole(ctx context.Context, clientID uuid.UUID, roleID uuid.UUID, appID int) ([]uuid.UUID, error) {
	const op = "service.ListPermissionsForRole"
	logger := s.logger.With(slog.String("op", op), slog.String("clientID", clientID.String()), slog.String("roleID", roleID.String()))
	logger.Debug("processing request")

	if err := utils.ValidateUUID(clientID); err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(roleID); err != nil {
		logger.Warn("invalid role_id format", slog.Any("error", err))
		return nil, fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	role, err := s.roleProvider.GetByID(ctx, clientID, roleID, appID)
	if err != nil {
		if errors.Is(err, serviceRole.ErrNotFound) {
			logger.Warn("role not found")
			return nil, ErrRoleNotFound
		}
		logger.Error("failed to get role", slog.Any("error", err))
		return nil, ErrInternal
	}

	if !role.IsActive {
		logger.Warn("attempt to work with inactive role")
		return nil, fmt.Errorf("%w: role is inactive", ErrInactiveEntity)
	}

	permissionIDs, err := s.relProvider.ListPermissionsForRole(ctx, roleID)
	if err != nil {
		logger.Error("failed to list permissions", slog.Any("error", err))
		return nil, ErrInternal
	}

	logger.Info("successfully retrieved permissions", slog.Int("count", len(permissionIDs)))
	return permissionIDs, nil
}

package service

import (
	servicePermission "backend/service/sso/permissions/service"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) ListRolesForPermission(ctx context.Context, clientID uuid.UUID, permissionID uuid.UUID, appID int) ([]uuid.UUID, error) {
	const op = "service.ListRolesForPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("clientID", clientID.String()), slog.String("permissionID", permissionID.String()), slog.Int("app_id", appID))
	logger.Debug("processing request")

	if err := utils.ValidateUUID(clientID); err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(permissionID); err != nil {
		logger.Warn("invalid permission_id format", slog.Any("error", err))
		return nil, fmt.Errorf("%w: permission_id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	perm, err := s.permProvider.GetByID(ctx, permissionID, appID)
	if err != nil {
		if errors.Is(err, servicePermission.ErrNotFound) {
			logger.Warn("permission not found")
			return nil, ErrPermissionNotFound
		}
		logger.Error("failed to get permission", slog.Any("error", err))
		return nil, ErrInternal
	}

	if !perm.IsActive {
		logger.Warn("attempt to work with inactive permission")
		return nil, fmt.Errorf("%w: permission is inactive", ErrInactiveEntity)
	}

	roles, err := s.relProvider.ListRolesForPermission(ctx, permissionID, clientID, appID)
	if err != nil {
		logger.Error("failed to list roles", slog.Any("error", err))
		return nil, ErrInternal
	}

	logger.Info("successfully retrieved roles", slog.Int("count", len(roles)))
	return roles, nil
}

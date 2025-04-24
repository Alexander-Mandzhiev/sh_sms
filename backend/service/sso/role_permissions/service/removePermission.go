package service

import (
	"backend/service/sso/models"
	"backend/service/sso/roles/service"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *service.Service) RemovePermission(ctx context.Context, clientID, roleID, permissionID uuid.UUID) (*models.Role, error) {
	const op = "service.Roles.RemovePermission"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", roleID.String()), slog.String("permission_id", permissionID.String()))
	logger.Debug("attempting to get user")
	return nil, nil
}

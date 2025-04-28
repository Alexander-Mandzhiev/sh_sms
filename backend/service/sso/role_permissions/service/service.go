package service

import (
	servicePermission "backend/service/sso/permissions/service"
	serviceRole "backend/service/sso/roles/service"
	"context"
	"errors"
	"github.com/google/uuid"

	"log/slog"
)

var (
	ErrInactiveEntity     = errors.New("inactive entity")
	ErrInternal           = errors.New("internal error")
	ErrPermissionNotFound = errors.New("permission not found")
	ErrRoleNotFound       = errors.New("role not found")
	ErrInvalidArgument    = errors.New("invalid argument")
)

type RolePermissionProvider interface {
	AddRolePermissions(ctx context.Context, roleID uuid.UUID, clientID uuid.UUID, appID int, permissionIDs []uuid.UUID) error
	RemoveRolePermissions(ctx context.Context, roleID uuid.UUID, clientID uuid.UUID, appID int, permissionIDs []uuid.UUID) error
	ListPermissionsForRole(ctx context.Context, roleID uuid.UUID, clientID uuid.UUID, appID int) ([]uuid.UUID, error)
	ListRolesForPermission(ctx context.Context, permissionID uuid.UUID, clientID uuid.UUID, appID int) ([]uuid.UUID, error)
	HasRelation(ctx context.Context, roleID, permissionID uuid.UUID, clientID uuid.UUID, appID int) (bool, error)
}

type Service struct {
	logger       *slog.Logger
	roleProvider serviceRole.RolesProvider
	permProvider servicePermission.PermissionProvider
	relProvider  RolePermissionProvider
}

func New(relProvider RolePermissionProvider, roleProvider serviceRole.RolesProvider, permProvider servicePermission.PermissionProvider, logger *slog.Logger) *Service {
	const op = "service.New"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing sso handle - service roles", slog.String("op", op))
	return &Service{
		relProvider:  relProvider,
		roleProvider: roleProvider,
		permProvider: permProvider,
		logger:       logger,
	}
}

package service

import (
	"log/slog"
)

type AuthService struct {
	users          UserService
	roles          RoleService
	permissions    PermissionService
	rolePermission RolePermissionService
	userRole       UserRoleService
	appManager     AppService
	clientApps     ClientAppsService
	secret         SecretService
	client         ClientService
	logger         *slog.Logger
}

func NewAuthService(userService UserService, roleService RoleService, permissionService PermissionService,
	rolePermission RolePermissionService, userRole UserRoleService, appManager AppService, secret SecretService,
	clientApps ClientAppsService, client ClientService, logger *slog.Logger) *AuthService {
	return &AuthService{
		users:          userService,
		roles:          roleService,
		permissions:    permissionService,
		userRole:       userRole,
		rolePermission: rolePermission,
		appManager:     appManager,
		clientApps:     clientApps,
		secret:         secret,
		client:         client,
		logger:         logger.With("service", "auth"),
	}
}

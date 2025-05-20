package service

import (
	config "backend/pkg/config/auth"
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
	session        SessionProvider
	logger         *slog.Logger
	cfg            *config.Config
}

func NewAuthService(userService UserService, roleService RoleService, permissionService PermissionService,
	rolePermission RolePermissionService, userRole UserRoleService, appManager AppService, secret SecretService,
	clientApps ClientAppsService, client ClientService, session SessionProvider,
	logger *slog.Logger, cfg *config.Config) *AuthService {
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
		session:        session,
		logger:         logger.With("service", "auth"),
		cfg:            cfg,
	}
}

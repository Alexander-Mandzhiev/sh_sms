package role_permissions_handle

import (
	"backend/protos/gen/go/sso/role_permissions"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type RolePermissionsService interface {
	AddPermissionsToRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error)
	RemovePermissionsFromRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error)
	ListPermissionsForRole(ctx context.Context, req *role_permissions.ListPermissionsRequest) (*role_permissions.ListPermissionsResponse, error)
	ListRolesForPermission(ctx context.Context, req *role_permissions.ListRolesRequest) (*role_permissions.ListRolesResponse, error)
	HasPermission(ctx context.Context, req *role_permissions.HasPermissionRequest) (*role_permissions.HasPermissionResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service RolePermissionsService
}

func New(service RolePermissionsService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	rp := router.Group("/role-permissions")
	{
		rp.POST("/add", h.addPermissionsToRole)
		rp.POST("/remove", h.removePermissionsFromRole)
		rp.POST("/list-permissions", h.listPermissionsForRole)
		rp.POST("/list-roles", h.listRolesForPermission)
		rp.POST("/check-permission", h.hasPermission)
	}
}

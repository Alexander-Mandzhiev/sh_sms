package user_roles_handle

import (
	"backend/protos/gen/go/sso/users_roles"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type UserRolesService interface {
	AssignRole(ctx context.Context, req *user_roles.AssignRequest) (*user_roles.UserRole, error)
	RevokeRole(ctx context.Context, req *user_roles.RevokeRequest) (*user_roles.RevokeResponse, error)
	ListForUser(ctx context.Context, req *user_roles.UserRequest) (*user_roles.UserRolesResponse, error)
	ListForRole(ctx context.Context, req *user_roles.RoleRequest) (*user_roles.UserRolesResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service UserRolesService
}

func New(service UserRolesService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	ur := router.Group("/user-roles")
	{
		ur.POST("/assign", h.assignRole)
		ur.POST("/revoke", h.revokeRole)
		ur.POST("/user", h.listForUser)
		ur.POST("/role", h.listForRole)
	}
}

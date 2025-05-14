package roles_handle

import (
	"backend/protos/gen/go/sso/roles"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type RolesService interface {
	CreateRole(ctx context.Context, req *roles.CreateRequest) (*roles.Role, error)
	GetRole(ctx context.Context, req *roles.GetRequest) (*roles.Role, error)
	UpdateRole(ctx context.Context, req *roles.UpdateRequest) (*roles.Role, error)
	DeleteRole(ctx context.Context, req *roles.DeleteRequest) (*roles.DeleteResponse, error)
	ListRoles(ctx context.Context, req *roles.ListRequest) (*roles.ListResponse, error)
	RestoreRole(ctx context.Context, req *roles.RestoreRequest) (*roles.Role, error)
}

type Handler struct {
	logger  *slog.Logger
	service RolesService
}

func New(service RolesService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {

	role := router.Group("/roles")
	{
		role.POST("", h.createRole)
		role.GET("/:id", h.getRole)
		role.PUT("/:id", h.updateRole)
		role.DELETE("/:id", h.deleteRole)
		role.POST("/list", h.listRoles)
		role.GET("/:id/restore", h.restoreRole)
	}
}

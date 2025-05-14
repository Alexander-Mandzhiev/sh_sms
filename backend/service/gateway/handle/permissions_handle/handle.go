package permissions_handle

import (
	"backend/protos/gen/go/sso/permissions"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type PermissionsService interface {
	CreatePermission(ctx context.Context, req *permissions.CreateRequest) (*permissions.Permission, error)
	GetPermission(ctx context.Context, req *permissions.GetRequest) (*permissions.Permission, error)
	UpdatePermission(ctx context.Context, req *permissions.UpdateRequest) (*permissions.Permission, error)
	DeletePermission(ctx context.Context, req *permissions.DeleteRequest) (*permissions.SuccessResponse, error)
	ListPermissions(ctx context.Context, req *permissions.ListRequest) (*permissions.ListResponse, error)
	RestorePermission(ctx context.Context, req *permissions.RestoreRequest) (*permissions.Permission, error)
}

type Handler struct {
	logger  *slog.Logger
	service PermissionsService
}

func New(service PermissionsService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {

	permission := router.Group("/permissions")
	{
		permission.POST("", h.create)
		permission.GET("/:id", h.get)
		permission.PUT("/:id", h.update)
		permission.DELETE("/:id", h.delete)
		permission.POST("/list", h.list)
		permission.GET("/:id/restore", h.restore)
	}
}

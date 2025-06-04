package classes_handle

import (
	library "backend/protos/gen/go/library"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type ClassesService interface {
	GetClass(ctx context.Context, req *library.GetClassRequest) (*library.Class, error)
	ListClasses(ctx context.Context) (*library.ListClassesResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service ClassesService
}

func New(service ClassesService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {

	permission := router.Group("/classes")
	{
		permission.GET("/", h.list)
		permission.GET("/:id", h.get)
	}
}

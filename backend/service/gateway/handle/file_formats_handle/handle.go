package file_formats_handle

import (
	library "backend/protos/gen/go/library"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type ClassesService interface {
	FileFormatExists(ctx context.Context, req *library.FileFormatExistsRequest) (*library.FileFormatExistsResponse, error)
	ListFileFormats(ctx context.Context) (*library.ListFileFormatsResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service ClassesService
}

func New(service ClassesService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {

	fileFormats := router.Group("/file_formats")
	{
		fileFormats.GET("/", h.list)
		fileFormats.GET("/:format", h.exist)
	}
}

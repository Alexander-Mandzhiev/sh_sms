package subjects_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type SubjectsService interface {
	CreateSubject(ctx context.Context, params *library_models.CreateSubjectParams) (*library.Subject, error)
	GetSubject(ctx context.Context, id int32) (*library.Subject, error)
	UpdateSubject(ctx context.Context, params *library_models.UpdateSubjectParams) (*library.Subject, error)
	DeleteSubject(ctx context.Context, id int32) error
	ListSubjects(ctx context.Context) (*library.ListSubjectsResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service SubjectsService
}

func New(service SubjectsService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	subjects := router.Group("/subjects")
	{
		subjects.POST("/", h.create)
		subjects.GET("/:id", h.get)
		subjects.PUT("/:id", h.update)
		subjects.DELETE("/:id", h.delete)
		subjects.GET("/", h.list)
	}
}

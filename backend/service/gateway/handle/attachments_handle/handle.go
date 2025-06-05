package attachments_handle

import (
	"backend/pkg/models/library"
	"backend/pkg/storage"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"mime/multipart"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, meta library_models.FileMetadata, file multipart.File) (*library_models.Attachment, error)
	GetAttachment(ctx context.Context, bookId int64, format string) (io.ReadSeekCloser, string, int64, error)
	DeleteAttachment(ctx context.Context, fileId string) error
	ListAttachmentsByBook(ctx context.Context, bookId int64) ([]*library_models.Attachment, error)
}

type Handler struct {
	logger  *slog.Logger
	storage storage.FileStorage
	service AttachmentsService
}

func New(service AttachmentsService, storage storage.FileStorage, logger *slog.Logger) *Handler {
	return &Handler{service: service, storage: storage, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	attachment := router.Group("/attachments")
	{
		attachment.POST("/", h.uploadFile)
		attachment.GET("/:book_id", h.listByBook)
		attachment.GET("/file/:book_id", h.get)
		attachment.DELETE("/:file_id", h.delete)
	}
}

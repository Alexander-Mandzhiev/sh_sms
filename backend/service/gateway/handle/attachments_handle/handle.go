package attachments_handle

import (
	"backend/pkg/models/library"
	"backend/pkg/storage"
	library "backend/protos/gen/go/library"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, req *library_models.CreateAttachmentRequest) (*library.Attachment, error)
	GetAttachment(ctx context.Context, bookId int64, format string) (*library.Attachment, error)
	DeleteAttachment(ctx context.Context, bookId int64, format string) (*emptypb.Empty, error)
	ListAttachmentsByBook(ctx context.Context, bookId int64) (*library.ListAttachmentsResponse, error)
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
		attachment.POST("/upload", h.uploadFile)
		attachment.GET("/:book_id", h.get)
		attachment.DELETE("/:book_id", h.delete)
		attachment.GET("/books/:book_id", h.listByBook)
		attachment.GET("/file/:file_id", h.downloadFile)
		attachment.DELETE("/file/:file_id", h.deleteFile)
	}
}

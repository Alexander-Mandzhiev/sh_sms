package attachments_handle

import (
	library "backend/protos/gen/go/library"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log/slog"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, req *library.CreateAttachmentRequest) (*library.Attachment, error)
	GetAttachment(ctx context.Context, req *library.GetAttachmentRequest) (*library.Attachment, error)
	UpdateAttachment(ctx context.Context, req *library.UpdateAttachmentRequest) (*library.Attachment, error)
	DeleteAttachment(ctx context.Context, req *library.DeleteAttachmentRequest) (*emptypb.Empty, error)
	ListAttachmentsByBook(ctx context.Context, req *library.ListAttachmentsByBookRequest) (*library.ListAttachmentsResponse, error)
	RestoreAttachment(ctx context.Context, req *library.RestoreAttachmentRequest) (*library.Attachment, error)
	DeleteFile(ctx context.Context, req *library.DeleteFileRequest) (*emptypb.Empty, error)
	UploadFile(ctx context.Context, metadata *library.FileMetadata, file io.Reader) (*library.UploadFileResponse, error)
}

type Handler struct {
	logger  *slog.Logger
	service AttachmentsService
}

func New(service AttachmentsService, logger *slog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) InitRoutes(router *gin.RouterGroup) {
	attachment := router.Group("/attachments")
	{
		attachment.POST("/", h.create)
		attachment.GET("/:book_id", h.get)
		attachment.PUT("/update", h.update)
		attachment.DELETE("/:book_id", h.delete)
		attachment.POST("/restore/:book_id", h.restore)
		attachment.GET("/books/:book_id", h.listByBook)
		attachment.DELETE("/file/:file_url", h.deleteFile)
		attachment.POST("/upload", h.upload)
	}
}

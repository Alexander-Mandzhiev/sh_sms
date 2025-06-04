package attachment_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/grpc"
	"io"
	"log/slog"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, req *library_models.CreateAttachmentRequest) (*library_models.Attachment, error)
	GetAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error)
	UpdateAttachment(ctx context.Context, req *library_models.UpdateAttachmentRequest) (*library_models.Attachment, error)
	DeleteAttachment(ctx context.Context, bookID int64, format string) error
	RestoreAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error)
	ListAttachmentsByBook(ctx context.Context, bookID int64, includeDeleted bool) ([]*library_models.Attachment, error)
	UploadFile(ctx context.Context, metadata *library_models.FileMetadata, file io.Reader) (*library_models.UploadedFile, error)
	DeleteFile(ctx context.Context, fileURL string) error
}
type serverAPI struct {
	library.UnimplementedAttachmentServiceServer
	service AttachmentsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service AttachmentsService, logger *slog.Logger) {
	library.RegisterAttachmentServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "attachment"),
	})
}

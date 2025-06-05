package attachment_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, req *library_models.CreateAttachmentRequest) (*library_models.Attachment, error)
	GetAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error)
	DeleteAttachment(ctx context.Context, bookID int64, format string) error
	ListAttachmentsByBook(ctx context.Context, bookID int64) ([]*library_models.Attachment, error)
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

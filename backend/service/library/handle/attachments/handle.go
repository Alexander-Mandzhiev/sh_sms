package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type AttachmentsService interface {
	Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.Attachment, error)
	Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.Attachment, error)
	Update(ctx context.Context, req *attachments.UpdateAttachmentRequest) (*attachments.Attachment, error)
	Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) error
	List(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.ListAttachmentsResponse, error)
	UploadFile(stream attachments.AttachmentsService_UploadFileServer) error
	DownloadFile(req *attachments.GetAttachmentRequest, stream attachments.AttachmentsService_DownloadFileServer) error
}

type serverAPI struct {
	attachments.UnimplementedAttachmentsServiceServer
	service AttachmentsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service AttachmentsService, logger *slog.Logger) {
	attachments.RegisterAttachmentsServiceServer(gRPCServer, &serverAPI{service: service,
		logger: logger.With("module", "attachments"),
	})
}

package attachment_types_handle

import (
	"backend/protos/gen/go/library/attachment_types"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type AttachmentTypesService interface {
	Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentType, error)
	Get(ctx context.Context, req *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentType, error)
	Update(ctx context.Context, req *attachment_types.UpdateAttachmentTypeRequest) (*attachment_types.AttachmentType, error)
	Delete(ctx context.Context, req *attachment_types.DeleteAttachmentTypeRequest) error
	List(ctx context.Context, req *attachment_types.ListAttachmentTypesRequest) (*attachment_types.ListAttachmentTypesResponse, error)
}

type serverAPI struct {
	attachment_types.UnimplementedAttachmentTypesServiceServer
	service AttachmentTypesService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service AttachmentTypesService, logger *slog.Logger) {
	attachment_types.RegisterAttachmentTypesServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "attachment_types"),
	})
}

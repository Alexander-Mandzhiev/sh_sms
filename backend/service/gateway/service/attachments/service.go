package attachments_service

import (
	"backend/pkg/storage"
	library "backend/protos/gen/go/library"
	"log/slog"
)

type AttachmentsService struct {
	client  library.AttachmentServiceClient
	storage storage.FileStorage
	logger  *slog.Logger
}

func NewAttachmentsService(client library.AttachmentServiceClient, storage storage.FileStorage, logger *slog.Logger) *AttachmentsService {
	return &AttachmentsService{
		client:  client,
		storage: storage,
		logger:  logger.With("service", "attachments"),
	}
}

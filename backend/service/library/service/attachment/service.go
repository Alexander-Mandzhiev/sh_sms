package attachments_service

import (
	library_models "backend/pkg/models/library"
	"backend/pkg/storage"
	"context"
	"log/slog"
)

type AttachmentsProvider interface {
	CreateAttachment(ctx context.Context, attachment *library_models.Attachment) error
	GetAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error)
	UpdateAttachment(ctx context.Context, attachment *library_models.Attachment) error
	DeleteAttachment(ctx context.Context, bookID int64, format string) error
	RestoreAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error)
	ListByBook(ctx context.Context, bookID int64, includeDeleted bool) ([]*library_models.Attachment, error)
}

type Service struct {
	logger      *slog.Logger
	fileStorage storage.FileStorage
	provider    AttachmentsProvider
}

func New(provider AttachmentsProvider, fileStorage storage.FileStorage, logger *slog.Logger) *Service {
	const op = "service.New.Library.Attachments"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service attachment", slog.String("op", op))
	return &Service{
		provider:    provider,
		fileStorage: fileStorage,
		logger:      logger,
	}
}

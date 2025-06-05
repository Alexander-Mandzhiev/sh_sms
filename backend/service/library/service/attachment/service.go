package attachments_service

import (
	library_models "backend/pkg/models/library"
	"context"
	"log/slog"
)

type AttachmentsProvider interface {
	CreateAttachment(ctx context.Context, attachment *library_models.Attachment) error
	GetAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error)
	DeleteAttachment(ctx context.Context, bookID int64, format string) error
	ListByBook(ctx context.Context, bookID int64) ([]*library_models.Attachment, error)
}

type Service struct {
	logger   *slog.Logger
	provider AttachmentsProvider
}

func New(provider AttachmentsProvider, logger *slog.Logger) *Service {
	const op = "service.New.Library.Attachments"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service attachment", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}

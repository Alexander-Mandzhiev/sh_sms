package attachments_service

import (
	library_models "backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type AttachmentsService interface {
	CreateAttachment(ctx context.Context, req *library_models.CreateAttachmentRequest) (*library.Attachment, error)
	GetAttachment(ctx context.Context, bookId int64, format string) (*library.Attachment, error)
	DeleteAttachment(ctx context.Context, bookId int64, format string) (*emptypb.Empty, error)
	ListAttachmentsByBook(ctx context.Context, bookId int64) (*library.ListAttachmentsResponse, error)
}

type attachmentsService struct {
	client library.AttachmentServiceClient
	logger *slog.Logger
}

func NewAttachmentsService(client library.AttachmentServiceClient, logger *slog.Logger) AttachmentsService {
	return &attachmentsService{
		client: client,
		logger: logger.With("service", "attachments"),
	}
}

func (s *attachmentsService) CreateAttachment(ctx context.Context, req *library_models.CreateAttachmentRequest) (*library.Attachment, error) {
	s.logger.Debug("Creating attachment", "book_id", req.BookId, "format", req.Format)
	request, err := req.CreateAttachmentRequestToProto()
	if err != nil {
		s.logger.Error("Failed to create attachment request", "error", err)
		return nil, err
	}
	return s.client.CreateAttachment(ctx, request)
}

func (s *attachmentsService) GetAttachment(ctx context.Context, bookId int64, format string) (*library.Attachment, error) {
	s.logger.Debug("Getting attachment", "book_id", bookId, "format", format)
	return s.client.GetAttachment(ctx, &library.GetAttachmentRequest{BookId: bookId, Format: format})
}

func (s *attachmentsService) DeleteAttachment(ctx context.Context, bookId int64, format string) (*emptypb.Empty, error) {
	s.logger.Debug("Deleting attachment", "book_id", bookId, "format", format)
	return s.client.DeleteAttachment(ctx, &library.DeleteAttachmentRequest{BookId: bookId, Format: format})
}

func (s *attachmentsService) ListAttachmentsByBook(ctx context.Context, bookId int64) (*library.ListAttachmentsResponse, error) {
	s.logger.Debug("Listing attachments by book", "book_id", bookId)
	return s.client.ListAttachmentsByBook(ctx, &library.ListAttachmentsByBookRequest{BookId: bookId})
}

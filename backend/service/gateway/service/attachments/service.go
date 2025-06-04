package attachments_service

import (
	library "backend/protos/gen/go/library"
	"context"
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

func (s *attachmentsService) CreateAttachment(ctx context.Context, req *library.CreateAttachmentRequest) (*library.Attachment, error) {
	s.logger.Debug("Creating attachment", "book_id", req.BookId, "format", req.Format)
	return s.client.CreateAttachment(ctx, req)
}

func (s *attachmentsService) GetAttachment(ctx context.Context, req *library.GetAttachmentRequest) (*library.Attachment, error) {
	s.logger.Debug("Getting attachment", "book_id", req.BookId, "format", req.Format)
	return s.client.GetAttachment(ctx, req)
}

func (s *attachmentsService) UpdateAttachment(ctx context.Context, req *library.UpdateAttachmentRequest) (*library.Attachment, error) {
	s.logger.Debug("Updating attachment", "book_id", req.BookId, "format", req.Format)
	return s.client.UpdateAttachment(ctx, req)
}

func (s *attachmentsService) DeleteAttachment(ctx context.Context, req *library.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	s.logger.Debug("Deleting attachment", "book_id", req.BookId, "format", req.Format)
	return s.client.DeleteAttachment(ctx, req)
}

func (s *attachmentsService) ListAttachmentsByBook(ctx context.Context, req *library.ListAttachmentsByBookRequest) (*library.ListAttachmentsResponse, error) {
	s.logger.Debug("Listing attachments by book", "book_id", req.BookId, "include_deleted", req.IncludeDeleted)
	return s.client.ListAttachmentsByBook(ctx, req)
}

func (s *attachmentsService) RestoreAttachment(ctx context.Context, req *library.RestoreAttachmentRequest) (*library.Attachment, error) {
	s.logger.Debug("Restoring attachment", "book_id", req.BookId, "format", req.Format)
	return s.client.RestoreAttachment(ctx, req)
}

func (s *attachmentsService) DeleteFile(ctx context.Context, req *library.DeleteFileRequest) (*emptypb.Empty, error) {
	s.logger.Debug("Deleting file", "file_url", req.FileUrl)
	return s.client.DeleteFile(ctx, req)
}

package attachments_service

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"fmt"
	"log/slog"
)

func (s *AttachmentsService) ListAttachmentsByBook(ctx context.Context, bookId int64) ([]*library_models.Attachment, error) {
	const op = "service.Gateway.Attachments.List"
	logger := s.logger.With(slog.String("op", op), slog.Int64("book_id", bookId))
	logger.Debug("Getting attachment")

	list, err := s.client.ListAttachmentsByBook(ctx, &library.ListAttachmentsByBookRequest{BookId: bookId})
	if err != nil {
		logger.Error("Error while calling client.ListAttachmentsByBook()")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var attachments []*library_models.Attachment
	for _, attachment := range list.Attachments {
		att := library_models.AttachmentFromProto(attachment)
		attachments = append(attachments, att)
	}

	return attachments, err
}

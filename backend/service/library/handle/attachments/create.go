package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Create(ctx context.Context, req *attachments.CreateAttachmentRequest) (*attachments.Attachment, error) {
	s.logger.Info("Create request received", "attachment_type_id", req.AttachmentTypeId)
	resp, err := s.service.Create(ctx, req)
	if err != nil {
		s.logger.Error("Create failed", "error", err)
		return nil, status.Error(codes.Internal, "failed to create attachment")
	}
	s.logger.Info("Create succeeded", "id", resp.Id)
	return resp, nil
}

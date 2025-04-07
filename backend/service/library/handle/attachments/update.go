package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Update(ctx context.Context, req *attachments.UpdateAttachmentRequest) (*attachments.Attachment, error) {
	s.logger.Info("Update request received", "id", req.Id)
	resp, err := s.service.Update(ctx, req)
	if err != nil {
		s.logger.Error("Update failed", "error", err)
		return nil, status.Error(codes.Internal, "update failed")
	}
	return resp, nil
}

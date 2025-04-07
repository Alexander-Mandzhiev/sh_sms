package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Get(ctx context.Context, req *attachments.GetAttachmentRequest) (*attachments.Attachment, error) {
	s.logger.Info("Get request received", "id", req.Id)
	resp, err := s.service.Get(ctx, req)
	if err != nil {
		s.logger.Error("Get failed", "error", err)
		return nil, status.Error(codes.NotFound, "attachment not found")
	}
	return resp, nil
}

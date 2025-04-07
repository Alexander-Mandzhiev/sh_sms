package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Delete(ctx context.Context, req *attachments.DeleteAttachmentRequest) (*attachments.DeleteAttachmentResponse, error) {
	s.logger.Info("Delete request received", "id", req.Id)
	err := s.service.Delete(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "delete failed")
	}
	return &attachments.DeleteAttachmentResponse{Success: true}, nil
}

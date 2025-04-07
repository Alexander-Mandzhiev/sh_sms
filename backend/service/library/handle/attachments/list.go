package attachments_handle

import (
	"backend/protos/gen/go/library/attachments"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) List(ctx context.Context, req *attachments.ListAttachmentsRequest) (*attachments.ListAttachmentsResponse, error) {
	s.logger.Info("List request received", "filters", req)
	resp, err := s.service.List(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "list failed")
	}
	return resp, nil
}

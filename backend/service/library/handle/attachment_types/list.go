package attachment_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/attachment_types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *attachment_types.ListAttachmentTypesRequest) (*attachment_types.ListAttachmentTypesResponse, error) {
	s.logger.Debug("ListAttachmentTypes called", slog.Any("filters", req))

	resp, err := s.service.List(ctx, req)
	if err != nil {
		s.logger.Error("Failed to list attachment types", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "list failed: %v", err)
	}

	s.logger.Info("Attachment types listed", slog.Int("count", len(resp.Items)))
	return resp, nil
}

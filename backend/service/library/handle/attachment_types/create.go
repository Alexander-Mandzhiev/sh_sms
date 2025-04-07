package attachment_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/attachment_types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *attachment_types.CreateAttachmentTypeRequest) (*attachment_types.AttachmentType, error) {
	s.logger.Debug("CreateAttachmentType called", slog.Any("request", req))

	resp, err := s.service.Create(ctx, req)
	if err != nil {
		s.logger.Error("Failed to create attachment type", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "failed to create attachment type: %v", err)
	}

	s.logger.Info("Attachment type created", slog.Int("id", int(resp.Id)))
	return resp, nil
}

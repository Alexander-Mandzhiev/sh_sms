package attachment_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/attachment_types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *attachment_types.GetAttachmentTypeRequest) (*attachment_types.AttachmentType, error) {
	s.logger.Debug("GetAttachmentType called", slog.Int("id", int(req.Id)))

	resp, err := s.service.Get(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get attachment type", sl.Err(err, true))
		return nil, status.Errorf(codes.NotFound, "attachment type not found: %v", err)
	}

	s.logger.Info("Attachment type retrieved", slog.Int("id", int(resp.Id)))
	return resp, nil
}

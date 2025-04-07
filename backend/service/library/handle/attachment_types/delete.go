package attachment_types_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/attachment_types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Delete(ctx context.Context, req *attachment_types.DeleteAttachmentTypeRequest) (*attachment_types.DeleteAttachmentTypeResponse, error) {
	s.logger.Debug("DeleteAttachmentType called", slog.Int("id", int(req.Id)))

	err := s.service.Delete(ctx, req)
	if err != nil {
		s.logger.Error("Failed to delete attachment type", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "delete failed: %v", err)
	}

	s.logger.Info("Attachment type deleted", slog.Int("id", int(req.Id)))
	return &attachment_types.DeleteAttachmentTypeResponse{Success: true}, nil
}

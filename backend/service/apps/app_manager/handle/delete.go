package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	const op = "handler.Delete"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Delete request received", slog.Int("id", int(req.GetId())))

	if req.GetId() <= 0 {
		logger.Error("Invalid ID")
		return nil, status.Error(codes.InvalidArgument, "invalid application ID")
	}

	if err := s.service.Delete(ctx, req.GetId()); err != nil {
		logger.Error("Delete failed", slog.Int("id", int(req.GetId())), slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Info("Application deleted", slog.Int("id", int(req.GetId())))
	return &pb.DeleteResponse{Success: true}, nil
}

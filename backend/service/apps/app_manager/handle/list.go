package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "handler.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("List request received", slog.Int("page", int(req.GetPage())), slog.Int("count", int(req.GetCount())))

	if req.GetPage() <= 0 || req.GetCount() <= 0 {
		logger.Error("Invalid pagination", slog.Int("page", int(req.GetPage())), slog.Int("count", int(req.GetCount())))
		return nil, status.Error(codes.InvalidArgument, "invalid pagination parameters")
	}

	res, err := s.service.List(ctx, req)
	if err != nil {
		logger.Error("List failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Debug("List results", slog.Int("total", int(res.GetTotalCount())), slog.Int("returned", len(res.GetApps())))
	return res, nil
}

package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.App, error) {
	const op = "handler.Update"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Update request received", slog.Int("id", int(req.GetId())))

	if req.GetId() <= 0 {
		logger.Error("Invalid ID")
		return nil, status.Error(codes.InvalidArgument, "invalid application ID")
	}

	if req.GetName() == "" && req.GetDescription() == "" && req.GetCode() == "" {
		logger.Error("No fields to update")
		return nil, status.Error(codes.InvalidArgument, "no fields provided for update")
	}

	app, err := s.service.Update(ctx, req)
	if err != nil {
		logger.Error("Update failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Info("Application updated", slog.Int("id", int(app.GetId())))
	return app, nil
}

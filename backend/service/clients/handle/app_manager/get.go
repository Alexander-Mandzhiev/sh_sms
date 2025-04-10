package app_manager_handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *pb.GetRequest) (*pb.App, error) {
	const op = "handler.Get"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Request received", slog.Any("request", req))

	if err := validateIdentifier(logger, req.Id, req.Code); err != nil {
		return nil, convertError(err)
	}

	app, err := s.service.Get(ctx, req)
	if err != nil {
		logger.Error("Get failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Debug("Application found", slog.Int("id", int(app.Id)))
	return app, nil
}

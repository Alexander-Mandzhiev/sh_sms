package app_manager_handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *serverAPI) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	const op = "handler.Delete"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Request received", slog.Any("request", req))

	if err := validateIdentifier(logger, req.Id, req.Code); err != nil {
		return nil, convertError(err)
	}

	res, err := s.service.Delete(ctx, req)
	if err != nil {
		logger.Error("Delete failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Info("Application deleted", slog.Any("request", req))
	return res, nil
}

package classes_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) GetClass(ctx context.Context, req *library.GetClassRequest) (*library.Class, error) {
	const op = "grpc.Library.Classes.GetById"
	logger := s.logger.With(slog.String("op", op), slog.Int("app_id", int(req.GetId())))
	logger.Debug("Get class by id called")

	if req.GetId() <= 0 {
		logger.Warn("id validation failed", slog.Any("error", ErrInvalidId))
		return nil, s.convertError(ErrInvalidId)
	}

	class, err := s.service.Get(ctx, int(req.GetId()))
	if err != nil {
		logger.Error("Failed to get class", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoClass := library_models.ClassToProto(class)
	logger.Info("Class retrieved", slog.Int("id", int(protoClass.GetId())))
	return protoClass, nil
}

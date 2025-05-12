package handle

import (
	"backend/protos/gen/go/clients/client_types"
	"backend/service/clients/client_types/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) RestoreClientType(ctx context.Context, req *client_types.RestoreRequest) (*client_types.ClientType, error) {
	const op = "grpc.client_types.Restore"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	logger.Debug("processing client type restoration")

	if req.GetId() <= 0 {
		err := fmt.Errorf("%w: invalid client type ID", ErrInvalidArgument)
		logger.Warn("validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	restoredType, err := s.service.Restore(ctx, int(req.GetId()))
	if err != nil {
		logger.Error("client type restoration failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client type restored successfully", slog.Int("id", restoredType.ID), slog.Bool("is_active", restoredType.IsActive))
	return models.ToProto(restoredType), nil
}

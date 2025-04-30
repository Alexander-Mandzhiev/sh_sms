package handle

import (
	"backend/protos/gen/go/clients/client_types"
	"backend/service/clients/client_types/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *client_types.GetRequest) (*client_types.ClientType, error) {
	const op = "grpc.client_types.Get"
	logger := s.logger.With(slog.String("op", op), slog.Int("requested_id", int(req.GetId())))
	logger.Debug("processing get client type request")

	if req.GetId() <= 0 {
		err := fmt.Errorf("%w: invalid client type ID", ErrInvalidArgument)
		logger.Warn("validation failed", slog.String("error_type", "invalid_argument"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	clientType, err := s.service.Get(ctx, int(req.GetId()))
	if err != nil {
		logger.Error("failed to get client type", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client type retrieved successfully", slog.Int("id", clientType.ID), slog.String("code", clientType.Code))
	return models.ToProto(clientType), nil
}

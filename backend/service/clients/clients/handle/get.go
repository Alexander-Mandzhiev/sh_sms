package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/clients/clients"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) GetClient(ctx context.Context, req *clients.GetRequest) (*clients.Client, error) {
	const op = "grpc.clients.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetId()))
	logger.Debug("processing client request")

	id, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid UUID format", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: id", ErrInvalidArgument))
	}

	client, err := s.service.Get(ctx, id)
	if err != nil {
		logger.Error("get failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client retrieved successfully")
	return client.ToProto(), nil
}

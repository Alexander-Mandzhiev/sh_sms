package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/clients/clients"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) RestoreClient(ctx context.Context, req *clients.RestoreRequest) (*clients.Client, error) {
	const op = "grpc.clients.Restore"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetId()))
	logger.Debug("processing client restoration")

	id, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid UUID format", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: id", ErrInvalidArgument))
	}

	client, err := s.service.Restore(ctx, id)
	if err != nil {
		logger.Error("restore failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client restore successfully")
	return client.ToProto(), nil
}

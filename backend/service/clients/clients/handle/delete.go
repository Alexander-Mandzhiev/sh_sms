package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/clients/clients"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) DeleteClient(ctx context.Context, req *clients.DeleteRequest) (*clients.DeleteResponse, error) {
	const op = "grpc.clients.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("id", req.GetId()))
	logger.Debug("processing client deletion")

	id, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid UUID format", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: id", ErrInvalidArgument))
	}

	permanent := false
	if req.Permanent != nil {
		permanent = req.GetPermanent()
	}

	if err = s.service.Delete(ctx, id, permanent); err != nil {
		logger.Error("delete failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("deletion successful", slog.String("id", id.String()), slog.Bool("permanent", permanent))
	return &clients.DeleteResponse{Success: true}, nil
}

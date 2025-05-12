package handle

import (
	"backend/protos/gen/go/clients/client_types"
	"backend/service/clients/client_types/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) UpdateClientType(ctx context.Context, req *client_types.UpdateRequest) (*client_types.ClientType, error) {
	const op = "grpc.client_types.Update"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	logger.Debug("processing client type update")

	if req.GetId() <= 0 {
		err := fmt.Errorf("%w: invalid client type ID", ErrInvalidArgument)
		logger.Warn("validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	params := models.UpdateParams{
		ID:          req.GetId(),
		Code:        req.GetCode(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}

	updatedType, err := s.service.Update(ctx, &params)
	if err != nil {
		logger.Error("client type update failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client type updated successfully", slog.Int("id", updatedType.ID), slog.String("code", updatedType.Code))
	return models.ToProto(updatedType), nil
}

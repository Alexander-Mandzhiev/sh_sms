package handle

import (
	"backend/protos/gen/go/clients/clients"
	"backend/service/clients/clients/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *clients.CreateRequest) (*clients.Client, error) {
	const op = "grpc.clients.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("name", req.GetName()), slog.Int("type_id", int(req.GetTypeId())))
	logger.Debug("processing client creation")

	if err := utils.ValidateString(req.GetName(), 255); err != nil {
		logger.Warn("invalid name", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: name", ErrInvalidArgument))
	}

	if err := utils.ValidateString(req.GetDescription(), 10000); err != nil {
		logger.Warn("invalid description", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: description", ErrInvalidArgument))
	}

	if err := utils.ValidateWebsite(req.GetWebsite(), 255); err != nil {
		logger.Warn("invalid website", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: website", ErrInvalidArgument))
	}

	if req.GetTypeId() <= 0 {
		logger.Warn("invalid type_id", slog.Int("value", int(req.GetTypeId())))
		return nil, s.convertError(ErrInvalidArgument)
	}

	params := &models.CreateParams{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		TypeID:      int(req.GetTypeId()),
		Website:     req.GetWebsite(),
	}

	client, err := s.service.Create(ctx, params)
	if err != nil {
		logger.Error("create failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("create successful", slog.String("id", client.ID))
	return client.ToProto(), nil
}

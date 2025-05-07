package handle

import (
	"backend/protos/gen/go/clients/clients"
	"backend/service/clients/clients/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
	"math"
)

func (s *serverAPI) Update(ctx context.Context, req *clients.UpdateRequest) (*clients.Client, error) {
	const op = "grpc.clients.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetId()))
	logger.Debug("processing client update")

	id, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid UUID format", slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: id", ErrInvalidArgument))
	}

	if req.Name == nil && req.Description == nil && req.TypeId == nil && req.Website == nil {
		logger.Warn("no fields to update")
		return nil, s.convertError(fmt.Errorf("%w: at least one field must be provided", ErrInvalidArgument))
	}

	params := &models.UpdateParams{ID: id.String()}

	if req.Name != nil {
		if err = utils.ValidateString(req.GetName(), 255); err != nil {
			logger.Warn("invalid name", slog.String("error", err.Error()))
			return nil, s.convertError(fmt.Errorf("%w: name", ErrInvalidArgument))
		}
		params.Name = req.Name
	}

	if req.Description != nil {
		if err = utils.ValidateString(req.GetDescription(), 10000); err != nil {
			logger.Warn("invalid description", slog.String("error", err.Error()))
			return nil, s.convertError(fmt.Errorf("%w: description", ErrInvalidArgument))
		}
		params.Description = req.Description
	}

	if req.TypeId != nil {
		tid := int(req.GetTypeId())
		if tid <= 0 || tid > math.MaxInt32 {
			logger.Warn("invalid type_id", slog.Int("value", tid))
			return nil, s.convertError(fmt.Errorf("%w: type_id", ErrInvalidArgument))
		}
		params.TypeID = &tid
	}

	if req.Website != nil {
		if err = utils.ValidateWebsite(req.GetWebsite(), 255); err != nil {
			logger.Warn("invalid website", slog.String("error", err.Error()))
			return nil, s.convertError(fmt.Errorf("%w: website", ErrInvalidArgument))
		}
		params.Website = req.Website
	}

	client, err := s.service.Update(ctx, params)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client updated successfully", slog.String("client_id", id.String()))
	return client.ToProto(), nil
}

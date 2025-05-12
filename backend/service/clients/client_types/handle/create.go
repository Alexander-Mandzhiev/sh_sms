package handle

import (
	"backend/protos/gen/go/clients/client_types"
	"backend/service/clients/client_types/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) CreateClientType(ctx context.Context, req *client_types.CreateRequest) (*client_types.ClientType, error) {
	const op = "grpc.client_types.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("code", req.GetCode()), slog.String("name", req.GetName()))
	logger.Debug("processing client type creation")

	if err := validateCreateRequest(req); err != nil {
		logger.Warn("request validation failed", slog.String("error_type", "invalid_request"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if err := utils.ValidateString(req.GetCode(), 50); err != nil {
		logger.Warn("invalid code format", slog.String("field", "code"), slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: code", ErrInvalidArgument))
	}

	if err := utils.ValidateString(req.GetName(), 100); err != nil {
		logger.Warn("invalid name format", slog.String("field", "name"), slog.String("error", err.Error()))
		return nil, s.convertError(fmt.Errorf("%w: name", ErrInvalidArgument))
	}

	clientType := &models.CreateParams{
		Code:        req.GetCode(),
		Name:        req.GetName(),
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	createdType, err := s.service.Create(ctx, clientType)
	if err != nil {
		logger.Error("client type creation failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("client type created successfully", slog.Int("id", createdType.ID))
	return models.ToProto(createdType), nil
}

func validateCreateRequest(req *client_types.CreateRequest) error {
	if req.GetCode() == "" {
		return fmt.Errorf("%w: code is required", ErrInvalidArgument)
	}
	if req.GetName() == "" {
		return fmt.Errorf("%w: name is required", ErrInvalidArgument)
	}
	return nil
}

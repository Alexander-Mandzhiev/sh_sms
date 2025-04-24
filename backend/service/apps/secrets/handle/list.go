package handle

import (
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/models"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "grpc.handler.Secret.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("List secrets request received")

	if err := utils.ValidatePagination(int(req.GetPage()), int(req.GetCount())); err != nil {
		logger.Warn("pagination validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	filter := models.ListFilter{
		Page:  int(req.Page),
		Count: int(req.Count),
	}

	if f := req.Filter; f != nil {
		filter.ClientID = f.ClientId
		filter.SecretType = f.SecretType

		if f.AppId != nil {
			appID := int(*f.AppId)
			filter.AppID = &appID
		}

		if f.IsActive != nil {
			filter.IsActive = f.IsActive
		}
	}

	if filter.ClientID != nil {
		if err := utils.ValidateUUIDToString(*filter.ClientID); err != nil {
			logger.Warn("invalid client_id in filter", slog.Any("error", err))
			return nil, s.convertError(constants.ErrInvalidArgument)
		}
	}

	if filter.SecretType != nil && !utils.IsValidSecretType(*filter.SecretType) {
		logger.Warn("invalid secret_type in filter", slog.String("type", *filter.SecretType))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	secrets, total, err := s.service.List(ctx, filter)
	if err != nil {
		logger.Error("failed to list secrets", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("Secrets list retrieved successfully", slog.Int("total_count", total), slog.Int("returned_count", len(secrets)))
	return &pb.ListResponse{
		Secrets:    s.convertSecretsToPB(secrets),
		TotalCount: int32(total),
		Page:       req.Page,
		Count:      req.Count,
	}, nil
}

package handle

import (
	"backend/pkg/utils"
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/models"
	"backend/service/constants"
	"context"
	"log/slog"
)

func (s *serverAPI) ListRotations(ctx context.Context, req *pb.ListRequest) (*pb.ListRotationHistoryResponse, error) {
	const op = "grpc.handler.Secret.ListRotations"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("ListRotations request received")

	if err := utils.ValidatePagination(int(req.Page), int(req.Count)); err != nil {
		logger.Warn("invalid pagination parameters", slog.Int("page", int(req.Page)), slog.Int("count", int(req.Count)), slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	filter := models.ListFilter{
		Page:  int(req.Page),
		Count: int(req.Count),
	}

	if f := req.Filter; f != nil {
		if f.ClientId != nil {
			if err := utils.ValidateUUIDToString(*f.ClientId); err != nil {
				logger.Warn("invalid client_id in filter", slog.Any("error", err))
				return nil, s.convertError(constants.ErrInvalidArgument)
			}
			filter.ClientID = f.ClientId
		}

		if f.AppId != nil {
			if *f.AppId <= 0 {
				logger.Warn("invalid app_id in filter", slog.Int("app_id", int(*f.AppId)))
				return nil, s.convertError(constants.ErrInvalidArgument)
			}
			appID := int(*f.AppId)
			filter.AppID = &appID
		}

		if f.SecretType != nil {
			if !utils.IsValidSecretType(*f.SecretType) {
				logger.Warn("invalid secret_type in filter", slog.String("type", *f.SecretType))
				return nil, s.convertError(constants.ErrInvalidArgument)
			}
			filter.SecretType = f.SecretType
		}

		if f.RotatedBy != nil {
			if err := utils.ValidateUUIDToString(*f.RotatedBy); err != nil {
				logger.Warn("invalid rotated_by in filter", slog.Any("error", err))
				return nil, s.convertError(constants.ErrInvalidArgument)
			}
			filter.RotatedBy = f.RotatedBy
		}

		if f.RotatedAfter != nil && f.RotatedBefore != nil {
			if f.RotatedAfter.AsTime().After(f.RotatedBefore.AsTime()) {
				logger.Warn("invalid time range", slog.Time("rotated_after", f.RotatedAfter.AsTime()), slog.Time("rotated_before", f.RotatedBefore.AsTime()))
				return nil, s.convertError(constants.ErrInvalidArgument)
			}
		}

		if f.RotatedAfter != nil {
			rotatedAfter := f.RotatedAfter.AsTime()
			filter.RotatedAfter = &rotatedAfter
		}

		if f.RotatedBefore != nil {
			rotatedBefore := f.RotatedBefore.AsTime()
			filter.RotatedBefore = &rotatedBefore
		}

		if f.IsActive != nil {
			logger.Warn("is_active filter is not supported for rotation history")
			return nil, s.convertError(constants.ErrInvalidArgument)
		}
	}

	rotations, total, err := s.service.ListRotations(ctx, filter)
	if err != nil {
		logger.Error("failed to list rotations", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("List history secrets rotations retrieved successfully", slog.Int("total_count", total), slog.Int("returned_count", len(rotations)))
	return &pb.ListRotationHistoryResponse{
		Rotations:  convertRotationsToPB(rotations),
		TotalCount: int32(total),
		Page:       req.Page,
		Count:      req.Count,
	}, nil
}

package handle

import (
	"backend/pkg/utils"
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/constants"
	"context"
	"errors"
	"log/slog"
)

func (s *serverAPI) GetRotation(ctx context.Context, req *pb.GetRotationHistoryRequest) (*pb.RotationHistory, error) {
	const op = "grpc.handler.Secret.GetRotation"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(req.GetId())))
	logger.Debug("starting rotation history retrieval")

	if req.GetId() <= 0 {
		logger.Warn("invalid rotation ID", slog.Int("id", int(req.GetId())))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	history, err := s.service.GetRotation(ctx, int(req.GetId()))
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("rotation history not found", slog.Int("id", int(req.GetId())))
		} else {
			logger.Error("database operation failed", slog.Any("error", err), slog.Int("id", int(req.GetId())))
		}
		return nil, s.convertError(err)
	}

	if err = utils.ValidateRotationHistory(history); err != nil {
		logger.Error("invalid rotation history data", slog.Any("error", err), slog.Any("history", slog.Group("rotation", slog.String("client_id", history.ClientID), slog.Int("app_id", history.AppID), slog.String("secret_type", history.SecretType), slog.Time("rotated_at", history.RotatedAt))))
		return nil, s.convertError(constants.ErrInternal)
	}

	logger.Info("rotation history retrieved successfully", slog.String("client_id", history.ClientID), slog.Int("app_id", history.AppID), slog.String("secret_type", history.SecretType), slog.Time("rotated_at", history.RotatedAt))
	return convertRotationHistoryToPB(history), nil
}

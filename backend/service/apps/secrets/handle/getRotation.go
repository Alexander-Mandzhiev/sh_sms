package handle

import (
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/constants"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) GetRotation(ctx context.Context, req *pb.GetRotationHistoryRequest) (*pb.RotationHistory, error) {
	const op = "grpc.handler.Secret.GetRotation"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())), slog.String("secret_type", req.GetSecretType()))
	logger.Debug("GetRotation request received")

	if err := utils.ValidateClientID(req.GetClientId()); err != nil {
		logger.Warn("empty client_id")
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id", slog.Int("app_id", int(req.GetAppId())))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if !utils.IsValidSecretType(req.SecretType) {
		logger.Warn("invalid secret_type")
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if err := utils.ValidateClientID(req.GetClientId()); err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if !utils.IsValidSecretType(req.GetSecretType()) {
		logger.Warn("invalid secret_type", slog.String("type", req.GetSecretType()))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if req.RotatedAt == nil {
		logger.Warn("rotated_at is required")
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	rotatedAt := req.RotatedAt.AsTime()
	if rotatedAt.IsZero() {
		logger.Warn("invalid rotated_at timestamp")
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	history, err := s.service.GetRotation(ctx, req.GetClientId(), int(req.GetAppId()), req.GetSecretType(), rotatedAt)
	if err != nil {
		logger.Error("failed to get rotation history", slog.Any("error", err), slog.Time("rotated_at", rotatedAt))
		return nil, s.convertError(err)
	}

	logger.Info("getting secret rotation history retrieved successfully")
	return convertRotationHistoryToPB(history), nil
}

package handle

import (
	"backend/pkg/utils"
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/constants"
	"context"
	"log/slog"
)

func (s *serverAPI) GetSecret(ctx context.Context, req *pb.GetRequest) (*pb.Secret, error) {
	const op = "grpc.handler.Secret.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()),
		slog.Int("app_id", int(req.GetAppId())), slog.String("secret_type", req.GetSecretType()))
	logger.Debug("Get secret request received")

	if err := utils.ValidateUUIDToString(req.GetClientId()); err != nil {
		logger.Warn("client_id validation failed", slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("app_id validation failed", slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if !utils.IsValidSecretType(req.GetSecretType()) {
		logger.Warn("invalid secret type")
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	secret, err := s.service.Get(ctx, req.GetClientId(), int(req.GetAppId()), req.GetSecretType())
	if err != nil {
		logger.Error("failed to get secret", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if secret == nil {
		return nil, constants.ErrNotFound
	}

	logger.Info("Secret retrieved successfully")
	return s.convertSecretToPB(secret), nil
}

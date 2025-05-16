package handle

import (
	"backend/pkg/utils"
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/constants"
	"context"
	"log/slog"
)

func (s *serverAPI) DeleteSecret(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	const op = "grpc.handler.Secret.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()),
		slog.Int("app_id", int(req.GetAppId())), slog.String("secret_type", req.GetSecretType()))
	logger.Debug("Delete secret request received")

	if err := utils.ValidateUUIDToString(req.GetClientId()); err != nil {
		logger.Warn("client_id validation failed", slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("app_id validation failed", slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if !utils.IsValidSecretType(req.GetSecretType()) {
		logger.Warn("invalid secret type", slog.String("type", req.GetSecretType()))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	err := s.service.Delete(ctx, req.GetClientId(), int(req.GetAppId()), req.GetSecretType())
	if err != nil {
		logger.Error("failed to delete secret", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("Secret deleted successfully")
	return &pb.DeleteResponse{Success: true}, nil
}

package handle

import (
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/constants"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) Revoke(ctx context.Context, req *pb.RevokeRequest) (*pb.Secret, error) {
	const op = "grpc.handler.Secret.Revoke"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()),
		slog.Int("app_id", int(req.GetAppId())), slog.String("secret_type", req.GetSecretType()))
	logger.Debug("Revoke secret request received")

	if err := utils.ValidateClientID(req.GetClientId()); err != nil {
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

	revokedSecret, err := s.service.Revoke(ctx, req.GetClientId(), int(req.GetAppId()), req.GetSecretType())
	if err != nil {
		logger.Error("failed to revoke secret", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if revokedSecret.RevokedAt == nil {
		logger.Error("revocation failed - timestamp not set")
		return nil, s.convertError(constants.ErrInternal)
	}

	logger.Info("Secret revoked successfully", slog.Time("revoked_at", *revokedSecret.RevokedAt), slog.Int("version", revokedSecret.SecretVersion))
	return s.convertSecretToPB(revokedSecret), nil
}

package handle

import (
	"backend/pkg/utils"
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/models"
	"backend/service/constants"
	"context"
	"log/slog"
)

func (s *serverAPI) RotateSecret(ctx context.Context, req *pb.RotateRequest) (*pb.Secret, error) {
	const op = "grpc.handler.Secret.Rotate"
	logger := s.logger.With(slog.String("op", op),
		slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())),
		slog.String("secret_type", req.GetSecretType()), slog.String("rotated_by", req.GetRotatedBy()))
	logger.Debug("Rotate secret request received")

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

	if req.GetRotatedBy() != "" {
		if err := utils.ValidateUUIDToString(req.GetRotatedBy()); err != nil {
			logger.Warn("rotated_by validation failed", slog.Any("error", err))
			return nil, s.convertError(constants.ErrInvalidArgument)
		}
	}

	params := models.RotateSecretParams{
		ClientID:   req.GetClientId(),
		AppID:      int(req.GetAppId()),
		SecretType: req.GetSecretType(),
		RotatedBy:  req.GetRotatedBy(),
	}

	rotatedSecret, err := s.service.Rotate(ctx, params)
	if err != nil {
		logger.Error("failed to rotate secret", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("Secret rotated successfully", slog.Int("new_version", rotatedSecret.SecretVersion), slog.Bool("active", rotatedSecret.RevokedAt == nil))
	return s.convertSecretToPB(rotatedSecret), nil
}

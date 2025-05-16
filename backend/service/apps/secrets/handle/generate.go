package handle

import (
	"backend/pkg/utils"
	pb "backend/protos/gen/go/apps/secrets"
	"backend/service/apps/models"
	"backend/service/constants"
	"context"
	"log/slog"
)

func (s *serverAPI) Generate(ctx context.Context, req *pb.CreateRequest) (*pb.Secret, error) {
	const op = "grpc.handler.Secret.Generate"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()),
		slog.Int("app_id", int(req.GetAppId())), slog.String("type", req.SecretType))
	logger.Debug("Generate secret request received")

	if err := utils.ValidateUUIDToString(req.GetClientId()); err != nil {
		logger.Warn("client_id validation failed", slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("app_id validation failed", slog.Any("error", err))
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	if !utils.IsValidSecretType(req.SecretType) {
		logger.Warn("invalid secret type")
		return nil, s.convertError(constants.ErrInvalidArgument)
	}

	params := models.CreateSecretParams{
		ClientID:   req.GetClientId(),
		AppID:      int(req.GetAppId()),
		SecretType: req.GetSecretType(),
		Algorithm:  "bcrypt",
	}

	if req.Algorithm != nil {
		params.Algorithm = req.GetAlgorithm()
	}

	createdSecret, err := s.service.Generate(ctx, params)
	if err != nil {
		logger.Error("failed to generate secret", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("Secret generated successfully")
	return s.convertSecretToPB(createdSecret), nil
}

package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/users"
	"context"
	"log/slog"
)

func (s *serverAPI) SetPassword(ctx context.Context, req *users.SetPasswordRequest) (*users.SuccessResponse, error) {
	const op = "grpc.user.SetPassword"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to set password")

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Log(ctx, slog.LevelWarn, "invalid client ID", slog.Any("error", err), slog.String("input", req.GetClientId()))
		return nil, s.convertError(err)
	}

	userID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Log(ctx, slog.LevelWarn, "invalid user ID", slog.Any("error", err), slog.String("input", req.GetId()))
		return nil, s.convertError(err)
	}

	if err = utils.ValidatePassword(req.NewPassword); err != nil {
		logger.Warn("password validation failed", slog.Any("error", err), slog.Int("password_length", len(req.NewPassword)))
		return nil, s.convertError(err)
	}

	logger.Debug("password validation passed", slog.Int("password_length", len(req.NewPassword)))

	if err = s.service.SetPassword(ctx, clientID, userID, req.NewPassword); err != nil {
		logger.Error("password update failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("password successfully updated")
	return &users.SuccessResponse{Success: true}, nil
}

package handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) Restore(ctx context.Context, req *users.RestoreRequest) (*users.User, error) {
	const op = "grpc.user.Restore"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to restore user")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	userID, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	user, err := s.service.Restore(ctx, clientID, userID)
	if err != nil {
		logger.Error("failed to restore user", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user retrieved successfully")
	return convertUserToProto(*user), nil
}

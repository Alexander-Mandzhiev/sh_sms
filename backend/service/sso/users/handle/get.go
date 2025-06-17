package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
	"context"
	"log/slog"
)

func (s *serverAPI) GetUser(ctx context.Context, req *users.GetRequest) (*users.User, error) {
	const op = "grpc.user.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to get user")

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	userID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	user, err := s.service.Get(ctx, clientID, userID)
	if err != nil {
		logger.Error("failed to get user", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user retrieved successfully")
	return models.ConvertUserToProto(user), nil
}

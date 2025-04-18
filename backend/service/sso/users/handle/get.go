package handle

import (
	"backend/protos/gen/go/sso/users"
	"backend/service/constants"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *users.GetRequest) (*users.User, error) {
	const op = "grpc.user.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to get user")

	clientID, err := uuid.Parse(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: invalid client_id", constants.ErrInvalidArgument))
	}

	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		logger.Warn("invalid user_id format", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: invalid user_id", constants.ErrInvalidArgument))
	}

	user, err := s.service.Get(ctx, clientID, userID)
	if err != nil {
		logger.Error("failed to get user", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user retrieved successfully")
	return convertUserToProto(*user), nil
}

package handle

import (
	"backend/pkg/utils"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"

	"backend/protos/gen/go/sso/users"
)

func (s *serverAPI) GetUserByLogin(ctx context.Context, req *users.GetUserByLoginRequest) (*users.UserInfo, error) {
	const op = "grpc.user.GetUserByLogin"
	logger := s.logger.With(slog.String("op", op), slog.String("email", req.Login))
	logger.Debug("attempting to create user")

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if err = utils.ValidateString(req.GetLogin(), 255); err != nil {
		logger.Warn("login validation failed", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: %v", ErrInvalidArgument, err))
	}

	if err = utils.ValidatePassword(req.GetPassword()); err != nil {
		logger.Warn("password validation failed", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: %v", ErrInvalidArgument, err))
	}

	userInfo, err := s.service.GetUserByLogin(ctx, req.GetLogin(), req.GetPassword(), clientID)
	if err != nil {
		logger.Error("create user failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user created successfully", slog.String("user_id", userInfo.ID.String()))
	return models.ConvertUserInfoToProto(userInfo), nil
}

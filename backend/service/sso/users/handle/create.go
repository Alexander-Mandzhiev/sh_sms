package handle

import (
	"context"
	"fmt"
	"log/slog"

	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
	"backend/service/utils"
)

func (s *serverAPI) CreateUser(ctx context.Context, req *users.CreateRequest) (*users.User, error) {
	const op = "grpc.user.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("email", req.Email))
	logger.Debug("attempting to create user")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	if err = utils.ValidateEmail(req.GetEmail()); err != nil {
		logger.Warn("email validation failed", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: %v", ErrInvalidArgument, err))
	}

	if err = utils.ValidatePassword(req.GetPassword()); err != nil {
		logger.Warn("password validation failed", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: %v", ErrInvalidArgument, err))
	}

	user := &models.User{
		ClientID: clientID,
		Email:    req.Email,
		FullName: req.FullName,
		Phone:    req.Phone,
	}

	if err = s.service.Create(ctx, user, req.Password); err != nil {
		logger.Error("create user failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user created successfully", slog.String("user_id", user.ID.String()))
	return convertUserToProto(*user), nil
}

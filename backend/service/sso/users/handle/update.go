package handle

import (
	"context"
	"fmt"
	"log/slog"

	"backend/protos/gen/go/sso/users"
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/utils"

	"github.com/google/uuid"
)

func (s *serverAPI) Update(ctx context.Context, req *users.UpdateRequest) (*users.User, error) {
	const op = "grpc.user.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to update user")

	clientID, err := uuid.Parse(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format")
		return nil, s.convertError(fmt.Errorf("%w: invalid client_id", constants.ErrInvalidArgument))
	}

	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		logger.Warn("invalid user_id format")
		return nil, s.convertError(fmt.Errorf("%w: invalid user_id", constants.ErrInvalidArgument))
	}

	//Добавить валидацию?

	updateData := &models.User{
		ID:       userID,
		ClientID: clientID,
	}

	if req.Email != nil {
		if err := utils.ValidateEmail(req.GetEmail()); err != nil {
			logger.Warn("invalid email format", slog.Any("error", err))
			return nil, s.convertError(fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err))
		}
		updateData.Email = req.GetEmail()
	}

	if req.FullName != nil {
		updateData.FullName = req.GetFullName()
	}

	if req.Phone != nil {
		updateData.Phone = req.GetPhone()
	}

	if req.IsActive != nil {
		updateData = req.GetIsActive()
	}

	updatedUser, err := s.service.Update(ctx, updateData)
	if err != nil {
		logger.Error("failed to update user", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user updated successfully")
	return convertUserToProto(*updatedUser), nil
}

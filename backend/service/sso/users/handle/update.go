package handle

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"backend/protos/gen/go/sso/users"
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/utils"
)

func (s *serverAPI) Update(ctx context.Context, req *users.UpdateRequest) (*users.User, error) {
	const op = "grpc.user.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to update user")

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

	updateData := &models.User{
		ID:       userID,
		ClientID: clientID,
		IsActive: true,
	}

	if req.Email != nil {
		if err = utils.ValidateEmail(req.GetEmail()); err != nil {
			logger.Warn("invalid email format", slog.Any("error", err))
			return nil, s.convertError(fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err))
		}
		updateData.Email = req.GetEmail()
	}

	if req.FullName != nil {
		if strings.TrimSpace(req.GetFullName()) == "" {
			logger.Warn("empty full name")
			return nil, s.convertError(fmt.Errorf("%w: full_name cannot be empty", constants.ErrInvalidArgument))
		}
		updateData.FullName = req.GetFullName()
	}

	if req.Phone != nil {
		if err = utils.ValidatePhone(req.GetPhone()); err != nil {
			logger.Warn("invalid phone format", slog.Any("error", err))
			return nil, s.convertError(fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err))
		}
		updateData.Phone = req.GetPhone()
	}

	if req.IsActive != nil {
		updateData.IsActive = req.GetIsActive()
		if !updateData.IsActive {
			logger = logger.With(slog.Bool("is_active", false))
			logger.Info("user deactivation requested")
		}
	}

	updatedUser, err := s.service.Update(ctx, updateData)
	if err != nil {
		logger.Error("failed to update user", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user updated successfully")
	return convertUserToProto(*updatedUser), nil
}

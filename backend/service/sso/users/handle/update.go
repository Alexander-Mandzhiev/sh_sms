package handle

import (
	"backend/pkg/utils"
	"context"
	"fmt"
	"log/slog"
	"strings"

	"backend/protos/gen/go/sso/users"
	"backend/service/sso/models"
)

func (s *serverAPI) UpdateUser(ctx context.Context, req *users.UpdateRequest) (*users.User, error) {
	const op = "grpc.user.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to update user")

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

	updateData := &models.User{
		ID:       userID,
		ClientID: clientID,
	}

	if req.Email != nil {
		if err = utils.ValidateEmail(req.GetEmail()); err != nil {
			logger.Warn("invalid email format", slog.Any("error", err))
			return nil, s.convertError(fmt.Errorf("%w: %v", ErrInvalidArgument, err))
		}
		updateData.Email = req.GetEmail()
	}

	if req.FullName != nil {
		if strings.TrimSpace(req.GetFullName()) == "" {
			logger.Warn("empty full name")
			return nil, s.convertError(fmt.Errorf("%w: full_name cannot be empty", ErrInvalidArgument))
		}
		updateData.FullName = req.GetFullName()
	}

	if req.Phone != nil {
		if req.GetPhone() != "" && !utils.IsValidPhone(req.GetPhone()) {
			logger.Warn("invalid phone format", slog.Any("error", err))
			return nil, s.convertError(fmt.Errorf("%w: %v", ErrInvalidArgument, err))
		}
		updateData.Phone = req.GetPhone()
	}

	updatedUser, err := s.service.Update(ctx, updateData)
	if err != nil {
		logger.Error("failed to update user", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("user updated successfully")
	return models.ConvertUserToProto(updatedUser), nil
}

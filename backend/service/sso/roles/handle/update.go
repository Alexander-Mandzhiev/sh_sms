package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/roles"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) UpdateRole(ctx context.Context, req *roles.UpdateRequest) (*roles.Role, error) {
	const op = "grpc.roles.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to update role")

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", ErrInvalidArgument))
	}

	roleID, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: role_id", ErrInvalidArgument))
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	updateData := models.Role{
		ID:       roleID,
		ClientID: clientID,
		AppID:    int(req.GetAppId()),
	}

	if req.Name != nil {
		if err = utils.ValidateRoleName(req.GetName(), 150); err != nil {
			logger.Warn("invalid role name", slog.Any("error", err))
			return nil, s.convertError(fmt.Errorf("%w: name cannot be empty", ErrInvalidArgument))
		}
		updateData.Name = *req.Name
	}

	if req.Description != nil {
		updateData.Description = *req.Description
	}

	if req.Level != nil {
		if err = utils.ValidateRoleLevel(int(req.GetLevel())); err != nil {
			return nil, s.convertError(fmt.Errorf("%w: invalid level", ErrInvalidArgument))
		}
		updateData.Level = int(*req.Level)
	}

	updatedRole, err := s.service.Update(ctx, &updateData)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("role updated successfully")
	return convertRoleToProto(updatedRole), nil
}

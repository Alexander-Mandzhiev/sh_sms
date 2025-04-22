package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *serverAPI) Update(ctx context.Context, req *roles.UpdateRequest) (*roles.Role, error) {
	const op = "grpc.roles.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to update role")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", constants.ErrInvalidArgument))
	}

	roleID, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: role_id", constants.ErrInvalidArgument))
	}

	updateData := models.Role{
		ID:       roleID,
		ClientID: clientID,
	}

	if req.Name != nil {
		if req.GetName() == "" {
			return nil, s.convertError(fmt.Errorf("%w: name cannot be empty", constants.ErrInvalidArgument))
		}
		updateData.Name = *req.Name
	}

	if req.Description != nil {
		updateData.Description = *req.Description
	}

	if req.Level != nil {
		if req.GetLevel() < 0 {
			return nil, s.convertError(fmt.Errorf("%w: invalid level", constants.ErrInvalidArgument))
		}
		updateData.Level = int(*req.Level)
	}

	if req.IsActive != nil {
		updateData.DeletedAt = nil
	}

	if req.ParentRoleId != nil {
		var parentID uuid.UUID
		parentID, err = utils.ValidateAndReturnUUID(*req.ParentRoleId)
		if err != nil {
			return nil, s.convertError(fmt.Errorf("%w: parent_role_id", constants.ErrInvalidArgument))
		}
		updateData.ParentRoleID = &parentID
	}

	if req.IsCustom != nil {
		updateData.IsCustom = *req.IsCustom
	}

	updatedRole, err := s.service.Update(ctx, &updateData)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("role updated successfully")
	return convertRoleToProto(updatedRole), nil
}

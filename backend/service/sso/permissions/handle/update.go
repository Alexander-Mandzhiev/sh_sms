package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/permissions"
	"backend/service/sso/models"
	"context"
	"log/slog"
)

func (s *serverAPI) UpdatePermission(ctx context.Context, req *permissions.UpdateRequest) (*permissions.Permission, error) {
	const op = "grpc.Permission.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("id", req.GetId()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("attempting to update permission")

	id, err := utils.ValidateStringAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid UUID format")
		return nil, s.convertError(ErrInvalidUUID)
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id")
		return nil, s.convertError(err)
	}

	updateData := models.Permission{
		ID:    id,
		AppID: int(req.GetAppId()),
	}

	if req.Code != nil {
		if err = utils.ValidateString(*req.Code, 100); err != nil {
			logger.Warn("invalid code format")
			return nil, s.convertError(ErrInvalidArgument)
		}
		updateData.Code = *req.Code
	}

	if req.Description != nil {
		updateData.Description = *req.Description
	}

	if req.Category != nil {
		if err = utils.ValidateString(req.GetCategory(), 50); err != nil {
			logger.Warn("invalid category format")
			return nil, s.convertError(ErrInvalidArgument)
		}
		updateData.Category = *req.Category
	}

	if req.IsActive != nil {
		updateData.IsActive = *req.IsActive
	}

	resultPerm, err := s.service.Update(ctx, updateData)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	return convertPermissionToProto(resultPerm), nil
}

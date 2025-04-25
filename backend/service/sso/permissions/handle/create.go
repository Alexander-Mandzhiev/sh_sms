package handle

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *permissions.CreateRequest) (*permissions.Permission, error) {
	const op = "grpc.Permission.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("code", req.GetCode()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("attempting to create permission")

	if err := utils.ValidateString(req.GetCode(), 100); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	newPerm := models.Permission{
		Code:        req.GetCode(),
		Description: req.GetDescription(),
		Category:    req.GetCategory(),
		AppID:       int(req.GetAppId()),
		IsActive:    true,
	}

	createdPerm, err := s.service.Create(ctx, newPerm)
	if err != nil {
		logger.Error("create permission failed", slog.Any("error", err), slog.String("code", req.GetCode()))
		return nil, s.convertError(err)
	}

	protoPerm := convertPermissionToProto(createdPerm)
	logger.Info("permission successfully created", slog.String("id", protoPerm.Id), slog.String("code", protoPerm.Code))
	return protoPerm, nil
}

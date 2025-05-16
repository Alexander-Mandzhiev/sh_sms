package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/permissions"
	"context"
	"log/slog"
)

func (s *serverAPI) GetPermission(ctx context.Context, req *permissions.GetRequest) (*permissions.Permission, error) {
	const op = "grpc.Permission.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("id", req.GetId()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("attempting to get permission")

	id, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid UUID format", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidUUID)
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("app_id validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	perm, err := s.service.Get(ctx, id, int(req.GetAppId()))
	if err != nil {
		logger.Error("get permission failed", slog.Any("error", err), slog.String("id", req.GetId()))
		return nil, s.convertError(err)
	}

	protoPerm := convertPermissionToProto(perm)
	logger.Debug("permission retrieved successfully", slog.String("code", protoPerm.Code), slog.Int("app_id", int(protoPerm.AppId)))
	return protoPerm, nil
}

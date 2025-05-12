package handle

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/service/utils"
	"context"
	"log/slog"
	"time"
)

func (s *serverAPI) RestorePermission(ctx context.Context, req *permissions.RestoreRequest) (*permissions.Permission, error) {
	const op = "grpc.Permission.Restore"
	logger := s.logger.With(slog.String("op", op), slog.String("permission_id", req.GetId()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("starting restore operation")

	id, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("validation failed: invalid UUID format", slog.Any("error", err), slog.String("input_id", req.GetId()))
		return nil, s.convertError(ErrInvalidUUID)
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("validation failed: invalid app_id", slog.Int("received_app_id", int(req.GetAppId())), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Debug("calling service layer")

	restoredPerm, err := s.service.Restore(ctx, id, int(req.GetAppId()))
	if err != nil {
		logger.Error("restore operation failed", slog.Any("error", err), slog.Bool("permanent", false))
		return nil, s.convertError(err)
	}

	logger.Info("permission successfully restored", slog.String("new_code", restoredPerm.Code), slog.Bool("is_active", restoredPerm.IsActive), slog.Time("restored_at", time.Now()))
	return convertPermissionToProto(restoredPerm), nil
}

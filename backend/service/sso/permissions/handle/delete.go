package handle

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) DeletePermission(ctx context.Context, req *permissions.DeleteRequest) (*permissions.SuccessResponse, error) {
	const op = "grpc.Permission.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("id", req.GetId()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("attempting to delete permission")

	id, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid UUID format", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidUUID)
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	err = s.service.Delete(ctx, id, int(req.GetAppId()), req.GetPermanent())
	if err != nil {
		logger.Error("delete failed", slog.Any("error", err), slog.Bool("permanent", req.GetPermanent()))
		return nil, s.convertError(err)
	}

	logger.Info("permission deleted successfully", slog.Bool("permanent", req.GetPermanent()))
	return &permissions.SuccessResponse{Success: true}, nil
}

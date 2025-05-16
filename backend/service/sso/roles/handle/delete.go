package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/roles"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) DeleteRole(ctx context.Context, req *roles.DeleteRequest) (*roles.DeleteResponse, error) {
	const op = "grpc.roles.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to delete role")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", ErrInvalidArgument))
	}

	roleID, err := utils.ValidateAndReturnUUID(req.GetId())
	if err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: role_id", ErrInvalidArgument))
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	err = s.service.Delete(ctx, clientID, roleID, int(req.GetAppId()), req.GetPermanent())
	if err != nil {
		logger.Error("delete role failed", slog.Any("error", err), slog.Bool("permanent", req.GetPermanent()))
		return nil, s.convertError(err)
	}

	logger.Info("role deleted successfully", slog.Bool("permanent", req.GetPermanent()))
	return &roles.DeleteResponse{Success: true}, nil
}

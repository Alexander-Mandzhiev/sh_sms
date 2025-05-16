package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/roles"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) RestoreRole(ctx context.Context, req *roles.RestoreRequest) (*roles.Role, error) {
	const op = "grpc.roles.Restore"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to restore role")

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

	role, err := s.service.Restore(ctx, clientID, roleID, int(req.GetAppId()))
	if err != nil {
		logger.Error("get role failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Debug("role retrieved successfully")
	return convertRoleToProto(role), nil
}

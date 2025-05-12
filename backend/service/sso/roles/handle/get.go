package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) GetRole(ctx context.Context, req *roles.GetRequest) (*roles.Role, error) {
	const op = "grpc.roles.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to get role")

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

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	role, err := s.service.Get(ctx, clientID, roleID, int(req.GetAppId()))
	if err != nil {
		logger.Error("get role failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Debug("role retrieved successfully")
	return convertRoleToProto(role), nil
}

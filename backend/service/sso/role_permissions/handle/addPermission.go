package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) AddPermission(ctx context.Context, req *roles.PermissionRequest) (*roles.Role, error) {
	const op = "grpc.roles.AddPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetRoleId()), slog.String("permission_id", req.GetPermissionId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to add permission to role")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", constants.ErrInvalidArgument))
	}

	roleID, err := utils.ValidateAndReturnUUID(req.GetRoleId())
	if err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: role_id", constants.ErrInvalidArgument))
	}

	permissionID, err := utils.ValidateAndReturnUUID(req.GetPermissionId())
	if err != nil {
		logger.Warn("invalid permission_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: permission_id", constants.ErrInvalidArgument))
	}

	updatedRole, err := s.service.AddPermission(ctx, clientID, roleID, permissionID)
	if err != nil {
		logger.Error("failed to add permission", slog.Any("error", err), slog.Any("client_id", clientID), slog.Any("role_id", roleID), slog.Any("permission_id", permissionID))
		return nil, s.convertError(err)
	}

	logger.Info("permission successfully added to role", slog.String("role_id", roleID.String()), slog.String("permission_id", permissionID.String()))
	return convertRoleToProto(updatedRole), nil
}

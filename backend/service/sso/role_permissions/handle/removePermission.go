package handle

import (
	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/constants"
	"backend/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *serverAPI) RemovePermissionsFromRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error) {
	const op = "grpc.role_permissions.RemovePermissionsFromRole"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetRoleId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("attempting to remove permission from role")

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

	if len(req.GetPermissionIds()) == 0 {
		logger.Warn("empty permission_ids list")
		return nil, s.convertError(fmt.Errorf("%w: permission_ids", ErrInvalidArgument))
	}

	permissionIDs := make([]uuid.UUID, 0, len(req.GetPermissionIds()))
	for _, elem := range req.GetPermissionIds() {
		var id uuid.UUID
		id, err = utils.ValidateAndReturnUUID(elem)
		if err != nil {
			logger.Warn("invalid permission_id", err)
			return nil, s.convertError(fmt.Errorf("%w: permission_id", ErrInvalidArgument))
		}
		permissionIDs = append(permissionIDs, id)
	}

	result, err := s.service.RemovePermissionsFromRole(ctx, clientID, roleID, int(req.GetAppId()), permissionIDs)
	if err != nil {
		logger.Error("failed to remove permission", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("successfully removed permissions", slog.Int("count", len(permissionIDs)))

	return &role_permissions.OperationStatus{
		Success:   result.Success,
		Message:   result.Message,
		Timestamp: timestamppb.New(result.OperationTime),
	}, nil
}

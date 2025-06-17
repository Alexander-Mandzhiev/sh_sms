package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/role_permissions"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *serverAPI) AddPermissionsToRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error) {
	const op = "grpc.roles.AddPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetRoleId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("processing request")

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", ErrInvalidArgument))
	}

	roleID, err := utils.ValidateStringAndReturnUUID(req.GetRoleId())
	if err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: role_id", ErrInvalidArgument))
	}

	if len(req.GetPermissionIds()) == 0 {
		logger.Warn("empty permission_ids list")
		return nil, s.convertError(fmt.Errorf("%w: permission_ids", ErrInvalidArgument))
	}

	permissionIDs := make([]uuid.UUID, 0, len(req.GetPermissionIds()))
	for _, elem := range req.GetPermissionIds() {
		var id uuid.UUID
		id, err = utils.ValidateStringAndReturnUUID(elem)
		if err != nil {
			logger.Warn("invalid permission_id", err)
			return nil, s.convertError(fmt.Errorf("%w: permission_id", ErrInvalidArgument))
		}
		permissionIDs = append(permissionIDs, id)
	}

	result, err := s.service.AddPermissionsToRole(ctx, clientID, roleID, int(req.GetAppId()), permissionIDs)
	if err != nil {
		logger.Error("failed to add permission", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("successfully added permissions", slog.Int("count", len(permissionIDs)))
	return &role_permissions.OperationStatus{
		Success:   result.Success,
		Message:   result.Message,
		Timestamp: timestamppb.New(result.OperationTime),
	}, nil
}

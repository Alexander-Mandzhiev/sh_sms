package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/role_permissions"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *serverAPI) HasPermission(ctx context.Context, req *role_permissions.HasPermissionRequest) (*role_permissions.HasPermissionResponse, error) {
	const op = "grpc.role_permissions.HasPermission"
	logger := s.logger.With(slog.String("op", op))

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid client_id")
	}

	roleID, err := utils.ValidateAndReturnUUID(req.GetRoleId())
	if err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid role_id")
	}

	permissionID, err := utils.ValidateAndReturnUUID(req.GetPermissionId())
	if err != nil {
		logger.Warn("invalid permission_id", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid permission_id")
	}

	hasPermission, err := s.service.HasPermission(ctx, clientID, roleID, permissionID, int(req.GetAppId()))
	if err != nil {
		logger.Error("permission check failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Debug("permissions check", slog.Any("permission", hasPermission))
	return &role_permissions.HasPermissionResponse{
		HasPermission: hasPermission,
		CheckedAt:     timestamppb.Now(),
	}, nil
}

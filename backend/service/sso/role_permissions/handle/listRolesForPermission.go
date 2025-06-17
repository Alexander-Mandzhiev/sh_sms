package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/role_permissions"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) ListRolesForPermission(ctx context.Context, req *role_permissions.ListRolesRequest) (*role_permissions.ListRolesResponse, error) {
	const op = "grpc.role_permissions.ListRolesForPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("permission_id", req.GetPermissionId()),
		slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("starting request processing")

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err), slog.String("input", req.GetClientId()))
		return nil, status.Error(codes.InvalidArgument, "invalid client_id format")
	}

	permissionID, err := utils.ValidateStringAndReturnUUID(req.GetPermissionId())
	if err != nil {
		logger.Warn("invalid permission_id format", slog.Any("error", err), slog.String("input", req.GetPermissionId()))
		return nil, status.Error(codes.InvalidArgument, "invalid permission_id format")
	}

	appID := int(req.GetAppId())
	if err = utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err), slog.Int("input", appID))
		return nil, status.Error(codes.InvalidArgument, "app_id must be positive integer")
	}

	roleIDs, err := s.service.ListRolesForPermission(ctx, clientID, permissionID, appID)
	if err != nil {
		logger.Error("failed to list roles", slog.Any("error", err), slog.Int("app_id", appID))
		return nil, s.convertError(err)
	}

	logger.Debug("successfully retrieved roles", slog.Int("count", len(roleIDs)))
	return &role_permissions.ListRolesResponse{
		RoleIds: convertUUIDsToStrings(roleIDs),
	}, nil
}

package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/role_permissions"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) ListPermissionsForRole(ctx context.Context, req *role_permissions.ListPermissionsRequest) (*role_permissions.ListPermissionsResponse, error) {
	const op = "grpc.role_permissions.ListPermissionsForRole"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetRoleId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("starting request processing")

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err), slog.String("input", req.GetClientId()))
		return nil, status.Error(codes.InvalidArgument, "invalid client_id format")
	}

	roleID, err := utils.ValidateAndReturnUUID(req.GetRoleId())
	if err != nil {
		logger.Warn("invalid role_id format", slog.Any("error", err), slog.String("input", req.GetRoleId()))
		return nil, status.Error(codes.InvalidArgument, "invalid role_id format")
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, s.convertError(ErrInvalidArgument)
	}

	permissionIDs, err := s.service.ListPermissionsForRole(ctx, clientID, roleID, int(req.GetAppId()))
	if err != nil {
		logger.Error("failed to list permissions", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Debug("successfully retrieved permissions", slog.Int("count", len(permissionIDs)))
	return &role_permissions.ListPermissionsResponse{
		PermissionIds: convertUUIDsToStrings(permissionIDs),
	}, nil
}

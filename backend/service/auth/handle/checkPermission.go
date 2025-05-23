package handle

import (
	"backend/pkg/jwt_manager"
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (h *serverAPI) CheckPermission(ctx context.Context, req *auth.PermissionCheckRequest) (*auth.PermissionCheckResponse, error) {
	const op = "grpc.handler.CheckPermission"
	logger := h.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())), slog.String("resource", req.GetResource()))

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid client_id format")
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid app_id format")
	}

	token := req.GetToken()
	if token == "" {
		logger.Warn("empty token")
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	resource := req.GetResource()
	if resource == "" {
		logger.Debug("resource is empty, using default")
		resource = "web app"
	}

	logger.Debug("checking permission", slog.String("resource", resource), slog.String("token_prefix", jwt_manager.HashToken(token)))

	permission := req.GetPermission()
	if permission == "" {
		logger.Debug("permission is empty, using default")
		return nil, status.Error(codes.InvalidArgument, "permission is required")
	}

	allowed, missingRoles, missingPermissions, err := h.service.CheckPermission(ctx, clientID, int(req.GetAppId()), resource, token, permission)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	logger.Debug("permission check completed", slog.Bool("allowed", allowed), slog.Any("missing_roles", missingRoles), slog.Any("missing_permissions", missingPermissions))
	return &auth.PermissionCheckResponse{
		Allowed:            allowed,
		MissingRoles:       missingRoles,
		MissingPermissions: missingPermissions,
	}, nil
}

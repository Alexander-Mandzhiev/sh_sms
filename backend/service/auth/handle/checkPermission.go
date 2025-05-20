package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (h *serverAPI) CheckPermission(ctx context.Context, req *auth.PermissionCheckRequest) (*auth.PermissionCheckResponse, error) {
	const op = "grpc.handler.CheckPermission"
	logger := h.logger.With(slog.String("op", op), slog.String("user_id", req.GetUserId()), slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())), slog.String("permission", req.GetPermission()))

	userID, err := utils.ValidateAndReturnUUID(req.GetUserId())
	if err != nil {
		logger.Warn("invalid user_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid user_id format")
	}

	if req.GetClientId() == "" {
		logger.Warn("client_id is empty")
		return nil, status.Error(codes.InvalidArgument, "client_id is required")
	}

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid client_id format")
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid app_id format")
	}

	if req.GetPermission() == "" {
		logger.Warn("permission is empty")
		return nil, status.Error(codes.InvalidArgument, "permission is required")
	}

	var resource string
	if req.Resource != "" {
		resource = req.Resource
	} else {
		logger.Debug("resource is empty, using default")
	}

	logger.Debug("checking permission", slog.String("resource", resource))

	perm := &models.PermissionCheck{
		UserID:     userID,
		ClientID:   clientID,
		AppID:      int(req.GetAppId()),
		Permission: req.GetPermission(),
		Resource:   resource,
	}

	allowed, missingRoles, missingPermissions, err := h.service.CheckPermission(ctx, perm)
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

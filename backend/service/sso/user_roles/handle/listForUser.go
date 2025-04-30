package handle

import (
	user_roles "backend/protos/gen/go/sso/users_roles"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) ListForUser(ctx context.Context, req *user_roles.UserRequest) (*user_roles.UserRolesResponse, error) {
	const op = "grpc.user_roles.ListForUser"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetUserId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("processing list request")

	if err := validateUserRequest(req); err != nil {
		logger.Warn("request validation failed", slog.String("error_type", "invalid_request"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	userID, err := utils.ValidateAndReturnUUID(req.GetUserId())
	if err != nil {
		logger.Warn("invalid user_id format", slog.String("field", "user_id"), slog.String("value", req.GetUserId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: user_id", ErrInvalidArgument))
	}

	clientID, err := utils.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.String("field", "client_id"), slog.String("value", req.GetClientId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", ErrInvalidArgument))
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id format", slog.String("field", "assigned_by"), slog.Int("value", int(req.GetAppId())), slog.Any("error", err))
		return nil, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	filter := models.ListRequest{
		UserID:   &userID,
		ClientID: &clientID,
		AppID:    utils.GetIntPointer(int(req.GetAppId())),
		Page:     int(req.Page),
		Count:    int(req.Count),
	}

	roles, total, err := s.service.ListForUser(ctx, filter)
	if err != nil {
		logger.Error("list processing failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	protoRoles := make([]*user_roles.UserRole, 0, len(roles))
	for _, role := range roles {
		protoRoles = append(protoRoles, convertUserRoleToProto(role))
	}

	logger.Info("list retrieved successfully", slog.Int("count", len(protoRoles)), slog.Int("total", total))

	return &user_roles.UserRolesResponse{
		Assignments: protoRoles,
		TotalCount:  int32(total),
		CurrentPage: req.Page,
		AppId:       req.AppId,
	}, nil
}

func validateUserRequest(req *user_roles.UserRequest) error {
	if req.GetUserId() == "" {
		return fmt.Errorf("%w: user_id is required", ErrInvalidArgument)
	}
	if req.GetClientId() == "" {
		return fmt.Errorf("%w: client_id is required", ErrInvalidArgument)
	}
	if req.GetAppId() <= 0 {
		return fmt.Errorf("%w: invalid app_id", ErrInvalidArgument)
	}
	if req.Page < 0 || req.Count <= 0 || req.Count > 1000 {
		return fmt.Errorf("%w: invalid pagination", ErrInvalidArgument)
	}
	return nil
}

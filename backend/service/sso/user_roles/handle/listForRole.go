package handle

import (
	utils2 "backend/pkg/utils"
	user_roles "backend/protos/gen/go/sso/users_roles"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *serverAPI) ListForRole(ctx context.Context, req *user_roles.RoleRequest) (*user_roles.UserRolesResponse, error) {
	const op = "grpc.user_roles.ListForRole"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", req.GetRoleId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("processing role list request")

	if err := validateRoleRequest(req); err != nil {
		logger.Warn("request validation failed", slog.String("error_type", "invalid_request"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	roleID, err := utils2.ValidateAndReturnUUID(req.GetRoleId())
	if err != nil {
		logger.Warn("invalid role_id format", slog.String("field", "role_id"), slog.String("value", req.GetRoleId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: role_id", ErrInvalidArgument))
	}

	clientID, err := utils2.ValidateAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.String("field", "client_id"), slog.String("value", req.GetClientId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", ErrInvalidArgument))
	}

	if err = utils2.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id format", slog.String("field", "assigned_by"), slog.Int("value", int(req.GetAppId())), slog.Any("error", err))
		return nil, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	filter := models.ListRequest{
		RoleID:   &roleID,
		ClientID: &clientID,
		AppID:    utils2.GetIntPointer(int(req.GetAppId())),
		Page:     int(req.Page),
		Count:    int(req.Count),
	}

	users, total, err := s.service.ListForRole(ctx, filter)
	if err != nil {
		logger.Error("list processing failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	protoUsers := make([]*user_roles.UserRole, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, convertUserRoleToProto(user))
	}

	logger.Info("role assignments retrieved", slog.Int("count", len(protoUsers)), slog.Int("total", total))
	return &user_roles.UserRolesResponse{
		Assignments: protoUsers,
		TotalCount:  int32(total),
		CurrentPage: req.Page,
		AppId:       req.AppId,
	}, nil
}

func validateRoleRequest(req *user_roles.RoleRequest) error {
	if req.GetRoleId() == "" {
		return fmt.Errorf("%w: role_id is required", ErrInvalidArgument)
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

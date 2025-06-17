package handle

import (
	"backend/pkg/utils"
	user_roles "backend/protos/gen/go/sso/users_roles"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (s *serverAPI) Assign(ctx context.Context, req *user_roles.AssignRequest) (*user_roles.UserRole, error) {
	const op = "grpc.user_roles.Assign"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetUserId()), slog.String("role_id", req.GetRoleId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("processing assignment request")

	if err := validateAssignRequest(req); err != nil {
		logger.Warn("request validation failed", slog.String("error_type", "invalid_request"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	userID, err := utils.ValidateStringAndReturnUUID(req.GetUserId())
	if err != nil {
		logger.Warn("invalid user_id format", slog.String("field", "user_id"), slog.String("value", req.GetUserId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: user_id", ErrInvalidArgument))
	}

	roleID, err := utils.ValidateStringAndReturnUUID(req.GetRoleId())
	if err != nil {
		logger.Warn("invalid role_id format", slog.String("field", "role_id"), slog.String("value", req.GetRoleId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: role_id", ErrInvalidArgument))
	}

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.String("field", "client_id"), slog.String("value", req.GetClientId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: client_id", ErrInvalidArgument))
	}

	assignedBy, err := utils.ValidateStringAndReturnUUID(req.GetAssignedBy())
	if err != nil {
		logger.Warn("invalid assigned_by format", slog.String("field", "assigned_by"), slog.String("value", req.GetAssignedBy()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: assigned_by", ErrInvalidArgument))
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id format", slog.String("field", "assigned_by"), slog.Int("value", int(req.GetAppId())), slog.Any("error", err))
		return nil, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	var expiresAt *time.Time
	if req.ExpiresAt != nil {
		t := req.ExpiresAt.AsTime()
		if t.Before(time.Now()) {
			logger.Warn("invalid expiration time", slog.String("field", "expires_at"), slog.Time("value", t), slog.String("validation", "expiration_time_must_be_in_future"))
			return nil, s.convertError(ErrInvalidExpiration)
		}
		expiresAt = &t
	}

	logger = logger.With(slog.String("app_id", fmt.Sprintf("%d", req.GetAppId())), slog.String("assigned_by", req.GetAssignedBy()))

	role := &models.UserRole{
		UserID:     userID,
		RoleID:     roleID,
		ClientID:   clientID,
		AppID:      int(req.GetAppId()),
		AssignedBy: assignedBy,
		ExpiresAt:  expiresAt,
		AssignedAt: time.Now().UTC(),
	}

	assignedRole, err := s.service.Assign(ctx, role)
	if err != nil {
		logger.Error("assignment processing failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("role assigned successfully", slog.Time("assigned_at", assignedRole.AssignedAt), slog.Any("expires_at", assignedRole.ExpiresAt))
	return convertUserRoleToProto(assignedRole), nil
}

func validateAssignRequest(req *user_roles.AssignRequest) error {
	if req.GetUserId() == "" {
		return fmt.Errorf("%w: user_id is required", ErrInvalidArgument)
	}
	if req.GetRoleId() == "" {
		return fmt.Errorf("%w: role_id is required", ErrInvalidArgument)
	}
	if req.GetClientId() == "" {
		return fmt.Errorf("%w: client_id is required", ErrInvalidArgument)
	}
	if req.GetAssignedBy() == "" {
		return fmt.Errorf("%w: assigned_by is required", ErrInvalidArgument)
	}
	if req.GetAppId() <= 0 {
		return fmt.Errorf("%w: invalid app_id", ErrInvalidArgument)
	}
	return nil
}

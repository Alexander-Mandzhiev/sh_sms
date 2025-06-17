package handle

import (
	"backend/pkg/utils"
	user_roles "backend/protos/gen/go/sso/users_roles"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (s *serverAPI) Revoke(ctx context.Context, req *user_roles.RevokeRequest) (*user_roles.RevokeResponse, error) {
	const op = "grpc.user_roles.Revoke"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", req.GetUserId()), slog.String("role_id", req.GetRoleId()), slog.String("client_id", req.GetClientId()))
	logger.Debug("processing revoke request")

	if err := validateRevokeRequest(req); err != nil {
		logger.Warn("request validation failed", slog.String("error_type", "invalid_request"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	userID, err := utils.ValidateStringAndReturnUUID(req.GetUserId())
	if err != nil {
		logger.Warn("invalid user_id format", slog.String("field", "user_id"), slog.String("value", req.GetUserId()), slog.Any("error", err))
		return nil, s.convertError(fmt.Errorf("%w: user_id", ErrInvalidArgument))
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id format", slog.String("field", "assigned_by"), slog.Int("value", int(req.GetAppId())), slog.Any("error", err))
		return nil, fmt.Errorf("%w: app_id", ErrInvalidArgument)
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

	err = s.service.Revoke(ctx, userID, roleID, clientID, int(req.GetAppId()))
	if err != nil {
		logger.Error("revoke processing failed", slog.String("error_type", "service_error"), slog.Any("error", err))
		return nil, s.convertError(err)
	}

	logger.Info("role revoked successfully")
	return &user_roles.RevokeResponse{
		Success:   true,
		RevokedAt: timestamppb.New(time.Now().UTC()),
	}, nil
}

func validateRevokeRequest(req *user_roles.RevokeRequest) error {
	if req.GetUserId() == "" {
		return fmt.Errorf("%w: user_id is required", ErrInvalidArgument)
	}
	if req.GetRoleId() == "" {
		return fmt.Errorf("%w: role_id is required", ErrInvalidArgument)
	}
	if req.GetClientId() == "" {
		return fmt.Errorf("%w: client_id is required", ErrInvalidArgument)
	}
	if req.GetAppId() <= 0 {
		return fmt.Errorf("%w: invalid app_id", ErrInvalidArgument)
	}
	return nil
}

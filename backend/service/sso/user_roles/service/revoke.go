package service

import (
	"backend/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Revoke(ctx context.Context, userID uuid.UUID, roleID uuid.UUID, clientID uuid.UUID, appID int) error {
	const op = "service.user_roles.Revoke"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("role_id", roleID.String()), slog.String("client_id", clientID.String()), slog.Int("app_id", appID))
	logger.Debug("processing role revocation")

	if err := utils.ValidateUUID(userID); err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return fmt.Errorf("%w: user_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(roleID); err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(clientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Int("app_id", appID))
		return fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	userExists, err := s.userProvider.Exists(ctx, clientID, userID)
	if err != nil {
		logger.Error("user existence check failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !userExists {
		logger.Warn("user not found")
		return ErrUserNotFound
	}

	roleExists, err := s.roleProvider.RoleExistsByID(ctx, clientID, roleID, appID)
	if err != nil {
		logger.Error("role existence check failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !roleExists {
		logger.Warn("role not found")
		return ErrRoleNotFound
	}

	assignmentExists, err := s.provider.Exists(ctx, userID, roleID, clientID, appID)
	if err != nil {
		logger.Error("assignment check failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !assignmentExists {
		logger.Warn("assignment not found")
		return ErrAssignmentNotFound
	}

	if err = s.provider.Revoke(ctx, userID, roleID, clientID, appID); err != nil {
		logger.Error("revoke operation failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("role revoked successfully")
	return nil
}

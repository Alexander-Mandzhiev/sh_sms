package service

import (
	"backend/pkg/utils"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Assign(ctx context.Context, role *models.UserRole) (*models.UserRole, error) {
	const op = "service.user_roles.Assign"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", role.UserID.String()), slog.String("role_id", role.RoleID.String()))
	logger.Debug("processing role assignment")

	if err := utils.ValidateUUID(role.UserID); err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: user_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(role.RoleID); err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(role.ClientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(role.AppID); err != nil {
		logger.Warn("invalid app_id", slog.Int("app_id", role.AppID))
		return nil, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	userExists, err := s.userProvider.Exists(ctx, role.ClientID, role.UserID)
	if err != nil {
		logger.Error("user existence check failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !userExists {
		logger.Warn("user not found")
		return nil, ErrUserNotFound
	}

	roleExists, err := s.roleProvider.RoleExistsByID(ctx, role.ClientID, role.RoleID, role.AppID)
	if err != nil {
		logger.Error("role existence check failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !roleExists {
		logger.Warn("role not found")
		return nil, ErrRoleNotFound
	}

	assignmentExists, err := s.provider.Exists(ctx, role.UserID, role.RoleID, role.ClientID, role.AppID)
	if err != nil {
		logger.Error("assignment check failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if assignmentExists {
		logger.Warn("assignment already exists")
		return nil, ErrAssignmentExists
	}

	assignedRole, err := s.provider.Assign(ctx, role)
	if err != nil {
		logger.Error("assignment creation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("role assigned successfully", slog.Time("assigned_at", assignedRole.AssignedAt))
	return assignedRole, nil
}

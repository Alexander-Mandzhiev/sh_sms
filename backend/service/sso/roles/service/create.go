package service

import (
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (s *Service) Create(ctx context.Context, role *models.Role) error {
	const op = "service.Roles.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", role.ClientID.String()), slog.String("role_name", role.Name))
	logger.Debug("attempting to create role")

	if err := utils.ValidateRoleName(role.Name, 150); err != nil {
		logger.Warn("invalid role name", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	if err := utils.ValidateRoleLevel(role.Level); err != nil {
		logger.Warn("invalid role level", slog.Int("level", role.Level))
		return fmt.Errorf("%w: level cannot be negative", ErrInvalidArgument)
	}

	roleExists, err := s.provider.RoleExists(ctx, role.ClientID, role.AppID, role.Name)
	if err != nil {
		logger.Error("role existence check failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if roleExists {
		logger.Warn("role name conflict", slog.String("name", role.Name))
		return fmt.Errorf("%w: role '%s'", ErrConflict, role.Name)
	}

	if role.ID == uuid.Nil {
		role.ID = uuid.New()
		logger.Debug("generated new role ID", slog.String("id", role.ID.String()))
	}

	if role.CreatedAt.IsZero() {
		role.CreatedAt = time.Now().UTC()
		role.UpdatedAt = role.CreatedAt
	}

	role.IsActive = true

	if err = s.provider.Create(ctx, role); err != nil {
		if errors.Is(err, ErrConflict) {
			logger.Warn("database conflict", slog.Any("role", role))
			return fmt.Errorf("%w: %v", ErrAlreadyExists, role.Name)
		}
		logger.Error("database operation failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("role created successfully", slog.String("role_id", role.ID.String()), slog.Int("level", int(role.Level)))
	return nil
}

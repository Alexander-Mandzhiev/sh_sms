package service

import (
	"backend/service/constants"
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

	if err := utils.ValidateRoleName(role.Name); err != nil {
		logger.Warn("invalid role name", slog.Any("error", err))
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if role.Level < 0 {
		logger.Warn("invalid role level", slog.Int("level", role.Level))
		return fmt.Errorf("%w: level cannot be negative", constants.ErrInvalidArgument)
	}

	roleExists, err := s.provider.RoleExists(ctx, role.ClientID, role.Name)
	if err != nil {
		logger.Error("role existence check failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}
	if roleExists {
		logger.Warn("role name conflict", slog.String("name", role.Name))
		return fmt.Errorf("%w: role '%s'", constants.ErrConflict, role.Name)
	}

	if role.ID == uuid.Nil {
		role.ID = uuid.New()
	}
	role.CreatedAt = time.Now().UTC()
	role.UpdatedAt = role.CreatedAt

	if !role.IsActive {
		role.IsActive = true
	}

	if err = s.provider.Create(ctx, role); err != nil {
		if errors.Is(err, constants.ErrConflict) {
			logger.Warn("database conflict", slog.Any("role", role))
			return fmt.Errorf("%w: %v", constants.ErrAlreadyExists, role.Name)
		}
		logger.Error("database operation failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	logger.Info("role created successfully", slog.String("role_id", role.ID.String()), slog.Int("level", role.Level))
	return nil
}

package service

import (
	"backend/service/constants"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (s *Service) Delete(ctx context.Context, clientID, roleID uuid.UUID, permanent bool) error {
	const op = "service.Roles.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", roleID.String()), slog.Bool("hard_delete", permanent))
	logger.Debug("attempting to delete role")

	role, err := s.provider.GetByID(ctx, clientID, roleID)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("role not found")
			return fmt.Errorf("%w: role", constants.ErrNotFound)
		}
		logger.Error("fetch role failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	if role.ClientID != clientID {
		logger.Warn("client ID mismatch", slog.String("expected", clientID.String()), slog.String("actual", role.ClientID.String()))
		return fmt.Errorf("%w: role access denied", constants.ErrPermissionDenied)
	}

	if !role.IsCustom && !permanent {
		logger.Warn("attempt to soft-delete system role")
		return fmt.Errorf("%w: system roles can't be soft-deleted", constants.ErrPermissionDenied)
	}

	if !permanent {
		var hasDependencies bool
		hasDependencies, err = s.provider.HasDependencies(ctx, roleID)
		if err != nil {
			logger.Error("dependency check failed", slog.Any("error", err))
			return fmt.Errorf("%w: %v", constants.ErrInternal, err)
		}

		if hasDependencies {
			logger.Warn("role has active dependencies")
			return fmt.Errorf("%w: role is in use", constants.ErrConflict)
		}
	}

	if permanent {
		if err = s.provider.HardDelete(ctx, clientID, roleID); err != nil {
			logger.Error("hard delete failed", slog.Any("error", err))
			return fmt.Errorf("%w: %v", constants.ErrInternal, err)
		}
	} else {
		deletedAt := time.Now().UTC()
		role.DeletedAt = &deletedAt
		role.IsActive = false
		if _, err = s.provider.Update(ctx, role); err != nil {
			logger.Error("soft delete failed", slog.Any("error", err))
			return fmt.Errorf("%w: %v", constants.ErrInternal, err)
		}
	}

	logger.Info("role deleted successfully")
	return nil
}

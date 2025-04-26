package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, clientID, roleID uuid.UUID, permanent bool) error {
	const op = "service.Roles.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", roleID.String()), slog.Bool("hard_delete", permanent))
	logger.Debug("attempting to delete role")

	role, err := s.provider.GetByID(ctx, clientID, roleID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("role not found")
			return fmt.Errorf("%w: role", ErrNotFound)
		}
		logger.Error("fetch role failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	if role.ClientID != clientID {
		logger.Warn("client ID mismatch", slog.String("expected", clientID.String()), slog.String("actual", role.ClientID.String()))
		return fmt.Errorf("%w: role access denied", ErrPermissionDenied)
	}

	if !role.IsCustom && !permanent {
		logger.Warn("attempt to soft-delete system role")
		return fmt.Errorf("%w: system roles can't be soft-deleted", ErrPermissionDenied)
	}

	if !permanent {
		var hasDependencies bool
		hasDependencies, err = s.provider.HasDependencies(ctx, roleID)
		if err != nil {
			logger.Error("dependency check failed", slog.Any("error", err))
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}

		if hasDependencies {
			logger.Warn("role has active dependencies")
			return fmt.Errorf("%w: role is in use", ErrConflict)
		}
	}

	if permanent {
		if err = s.provider.HardDelete(ctx, clientID, roleID); err != nil {
			logger.Error("hard delete failed", slog.Any("error", err))
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	} else {
		if err = s.provider.SoftDelete(ctx, clientID, roleID); err != nil {
			logger.Error("soft delete failed", slog.Any("error", err))
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	logger.Info("role deleted successfully")
	return nil
}

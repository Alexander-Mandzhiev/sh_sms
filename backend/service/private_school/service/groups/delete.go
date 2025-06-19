package groups_service

import (
	"backend/pkg/utils"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) DeleteGroup(ctx context.Context, publicID, clientID uuid.UUID) error {
	const op = "service.PrivateSchool.Groups.DeleteGroup"
	logger := s.logger.With(slog.String("op", op), slog.String("public_id", publicID.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("deleting group")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return ctx.Err()
	}

	err := s.provider.DeleteGroup(ctx, publicID, clientID)
	if err != nil {
		return s.handleRepoError(err, op)
	}

	logger.Info("group deleted")
	return nil
}

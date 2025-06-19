package groups_service

import (
	groups_models "backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) GetGroup(ctx context.Context, publicID, clientID uuid.UUID) (*groups_models.Group, error) {
	const op = "service.PrivateSchool.Groups.GetGroup"
	logger := s.logger.With(slog.String("op", op), slog.String("public_id", publicID.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("getting group")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	group, err := s.provider.GetGroup(ctx, publicID, clientID)
	if err != nil {
		return nil, s.handleRepoError(err, op)
	}

	logger.Debug("group retrieved")
	return group, nil
}

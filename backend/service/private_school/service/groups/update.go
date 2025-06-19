package groups_service

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"log/slog"
)

func (s *Service) UpdateGroup(ctx context.Context, req *groups_models.UpdateGroup) (*groups_models.Group, error) {
	const op = "service.PrivateSchool.Groups.UpdateGroup"
	logger := s.logger.With(slog.String("op", op), slog.String("public_id", req.PublicID.String()), slog.String("client_id", req.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("updating group")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	group, err := s.provider.UpdateGroup(ctx, req)
	if err != nil {
		return nil, s.handleRepoError(err, op, "name", req.Name)
	}

	logger.Info("group updated", "public_id", group.PublicID)
	return group, nil
}

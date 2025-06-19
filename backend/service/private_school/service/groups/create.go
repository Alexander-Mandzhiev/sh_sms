package groups_service

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"log/slog"
)

func (s *Service) CreateGroup(ctx context.Context, req *groups_models.CreateGroup) (*groups_models.Group, error) {
	const op = "service.PrivateSchool.Groups.CreateGroup"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.ClientID.String()), slog.String("name", req.Name), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("creating group")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	group, err := s.provider.CreateGroup(ctx, req)
	if err != nil {
		return nil, s.handleRepoError(err, op, "name", req.Name)
	}

	logger.Info("group created", "public_id", group.PublicID)
	return group, nil
}

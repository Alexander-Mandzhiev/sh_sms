package groups_service

import (
	groups_models "backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"log/slog"
)

func (s *Service) ListGroups(ctx context.Context, req *groups_models.ListGroupsRequest) (*groups_models.GroupListResponse, error) {
	const op = "service.PrivateSchool.Groups.ListGroups"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.ClientID.String()), slog.Int("page_size", int(req.PageSize)), slog.String("trace_id", utils.TraceIDFromContext(ctx)))

	if req.Cursor != nil {
		logger = logger.With(slog.Int64("cursor", *req.Cursor))
	}
	if req.NameFilter != nil {
		logger = logger.With(slog.String("name_filter", *req.NameFilter))
	}

	logger.Debug("listing groups")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	response, err := s.provider.ListGroups(ctx, req)
	if err != nil {
		return nil, s.handleRepoError(err, op)
	}

	logger.Info("groups listed", "count", len(response.Groups), "has_next", response.NextCursor > 0)
	return response, nil
}

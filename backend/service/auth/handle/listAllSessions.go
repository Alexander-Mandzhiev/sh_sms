package handle

import (
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (h *serverAPI) ListAllSessions(ctx context.Context, req *auth.AllSessionsFilter) (*auth.SessionList, error) {
	const op = "grpc.handler.ListAllSessions"
	logger := h.logger.With(
		slog.String("op", op),
		slog.String("client_id", req.GetClientId()),
		slog.Int("app_id", int(req.GetAppId())),
		slog.Bool("active_only", req.GetActiveOnly()),
	)

	filter, err := models.AllSessionsFilterFromProto(req)
	if err != nil {
		logger.Warn("invalid filter format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid filter parameters")
	}

	logger.Debug("fetching all sessions", slog.Any("full_name", filter.FullName), slog.Any("email", filter.Email), slog.Any("phone", filter.Phone))

	sessions, err := h.service.ListAllSessions(ctx, *filter)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	logger.Info("all sessions fetched", slog.Int("count", len(sessions)))
	return models.SessionsToProto(sessions), nil
}

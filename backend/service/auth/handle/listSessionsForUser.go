package handle

import (
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (h *serverAPI) ListSessionsForUser(ctx context.Context, req *auth.SessionFilter) (*auth.SessionList, error) {
	const op = "grpc.handler.ListSessionsForUser"
	logger := h.logger.With(
		slog.String("op", op),
		slog.String("user_id", req.GetUserId()),
		slog.String("client_id", req.GetClientId()),
		slog.Int("app_id", int(req.GetAppId())),
		slog.Bool("active_only", req.GetActiveOnly()),
	)

	if req.GetUserId() == "" {
		logger.Warn("missing user_id in request")
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	filter, err := models.SessionFilterFromProto(req)
	if err != nil {
		logger.Warn("failed to parse filter", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid filter parameters")
	}

	logger.Debug("fetching user sessions", slog.Int("page", filter.Page), slog.Int("count", filter.Count))

	sessions, err := h.service.ListSessionsForUser(ctx, *filter)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	pbSessionList, err := models.SessionsToProto(sessions)
	if err != nil {
		logger.Error("failed to convert sessions to proto", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "failed to process sessions")
	}

	logger.Info("all sessions fetched", slog.Int("count", len(sessions)))
	return pbSessionList, nil
}

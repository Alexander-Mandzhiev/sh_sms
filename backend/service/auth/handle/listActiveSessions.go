package handle

import (
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (h *serverAPI) ListActiveSessions(ctx context.Context, req *auth.SessionFilter) (*auth.SessionList, error) {
	const op = "grpc.handler.ListActiveSessions"
	logger := h.logger.With(slog.String("op", op), slog.String("user_id", req.GetUserId()), slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())), slog.Bool("active_only", req.GetActiveOnly()))

	filter, err := models.SessionFilterFromProto(req)
	if err != nil {
		logger.Warn("failed to parse filter", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid filter")
	}

	logger.Debug("fetching sessions")

	sessions, err := h.service.ListActiveSessions(ctx, *filter)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	logger.Debug("sessions fetched successfully", slog.Int("count", len(sessions)))
	return models.SessionsToProto(sessions), nil
}

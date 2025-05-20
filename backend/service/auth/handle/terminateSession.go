package handle

import (
	"backend/protos/gen/go/auth"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (h *serverAPI) TerminateSession(ctx context.Context, req *auth.SessionID) (*emptypb.Empty, error) {
	const op = "grpc.handler.TerminateSession"
	sessionIDStr := req.GetSessionId()
	logger := h.logger.With(slog.String("op", op), slog.String("session_id", sessionIDStr))

	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		logger.Warn("invalid session_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid session_id format")
	}

	logger.Debug("terminating session")

	err = h.service.TerminateSession(ctx, sessionID)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	logger.Debug("session terminated successfully")
	return &emptypb.Empty{}, nil
}

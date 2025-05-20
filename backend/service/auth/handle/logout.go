package handle

import (
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/auth"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (h *serverAPI) Logout(ctx context.Context, req *auth.LogoutRequest) (*emptypb.Empty, error) {
	const op = "grpc.handler.Logout"
	logger := h.logger.With(slog.String("op", op))

	if req.AccessToken == "" {
		logger.Warn("access_token is empty")
		return nil, status.Error(codes.InvalidArgument, "access_token is required")
	}

	if req.RefreshToken == "" {
		logger.Warn("refresh_token is empty")
		return nil, status.Error(codes.InvalidArgument, "refresh_token is required")
	}

	logger.Debug("initiating logout", slog.String("access_token_hash", jwt_manager.HashToken(req.AccessToken)))

	err := h.service.Logout(ctx, req.AccessToken, req.RefreshToken)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	logger.Debug("user successfully logged out")
	return &emptypb.Empty{}, nil
}

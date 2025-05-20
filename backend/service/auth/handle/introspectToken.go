package handle

import (
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (h *serverAPI) IntrospectToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenIntrospection, error) {
	const op = "grpc.handler.IntrospectToken"
	logger := h.logger.With(slog.String("op", op), slog.String("token_type_hint", req.TokenTypeHint))

	if req.Token == "" {
		logger.Warn("token is empty")
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	tokenHash := jwt_manager.HashToken(req.Token)
	logger.Debug("introspecting token", slog.String("token_hash", tokenHash))

	validationResult, err := h.service.IntrospectToken(ctx, req.Token, req.TokenTypeHint)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	logger.Debug("token introspection completed", slog.String("token_hash", tokenHash))

	return models.TokenIntrospectionToProto(validationResult), nil
}

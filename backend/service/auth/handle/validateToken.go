package handle

import (
	"backend/pkg/jwt_manager"
	"backend/service/auth/models"
	"context"
	"log/slog"

	"backend/protos/gen/go/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *serverAPI) ValidateToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenInfo, error) {
	const op = "grpc.handler.ValidateToken"
	logger := h.logger.With(slog.String("op", op), slog.String("token_type_hint", req.TokenTypeHint))

	if req.Token == "" {
		logger.Warn("token is empty")
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	tokenHash := jwt_manager.HashToken(req.Token)
	logger.Debug("validating token", slog.String("token_hash", tokenHash))

	validationResult, err := h.service.ValidateToken(ctx, req.Token, req.TokenTypeHint)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	logger.Debug("token validated successfully", slog.String("token_hash", tokenHash))
	return models.TokenValidationToProto(validationResult), nil
}

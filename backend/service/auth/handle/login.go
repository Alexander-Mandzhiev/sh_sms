package handle

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (h *serverAPI) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	const op = "grpc.handler.Login"
	logger := h.logger.With(slog.String("op", op), slog.String("client_id", req.ClientId), slog.Int("app_id", int(req.AppId)))

	clientID, err := utils.ValidateStringAndReturnUUID(req.ClientId)
	if err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid client_id format")
	}

	if err = utils.ValidateAppID(int(req.AppId)); err != nil {
		logger.Warn("invalid app_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid app_id format")
	}

	authReq, err := models.AuthRequestFromProto(req)
	if err != nil {
		logger.Warn("failed to convert auth request", slog.Any("error", err))
		return nil, status.Errorf(codes.InvalidArgument, "%s: %v", op, err)
	}

	userInfo, accessToken, refreshToken, err := h.service.Login(ctx, *authReq)
	if err != nil {
		return nil, h.convertError(op, err)
	}

	metadata := &auth.TokenMetadata{
		ClientId:  clientID.String(),
		AppId:     req.AppId,
		TokenType: "Bearer",
		Issuer:    h.cfg.ServiceName,
		Audiences: h.cfg.JWT.Audiences,
	}

	logger.Debug("user successfully logged in", slog.String("user_id", userInfo.ID))

	return &auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(time.Now().Add(h.cfg.JWT.AccessDuration)),
		User:         models.ConvertUserToProto(userInfo),
		Metadata:     metadata,
	}, nil
}

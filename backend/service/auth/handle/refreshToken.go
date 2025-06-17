package handle

import (
	"backend/pkg/utils"
	"context"
	"log/slog"
	"time"

	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *serverAPI) RefreshToken(ctx context.Context, req *auth.RefreshRequest) (*auth.AuthResponse, error) {
	const op = "grpc.handler.RefreshToken"
	logger := h.logger.With(slog.String("op", op), slog.String("client_id", req.ClientId), slog.Int("app_id", int(req.AppId)))

	clientID, err := utils.ValidateStringAndReturnUUID(req.GetClientId())
	if err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid client_id format")
	}

	if err = utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("invalid app_id format", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "invalid app_id format")
	}

	userInfo, accessToken, refreshToken, err := h.service.RefreshToken(ctx, req.RefreshToken, clientID, int(req.AppId))
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

	return &auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(time.Now().Add(h.cfg.JWT.AccessDuration)),
		User:         models.ConvertUserToProto(userInfo),
		Metadata:     metadata,
	}, nil
}

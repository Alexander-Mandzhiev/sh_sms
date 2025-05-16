package handle

import (
	"backend/protos/gen/go/auth"
	"backend/service/auth/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (h *serverAPI) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	const op = "grpc.handler.Login"

	authReq, err := models.AuthRequestFromProto(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %v", op, err)
	}

	userInfo, accessToken, refreshToken, err := h.service.Login(ctx, *authReq)
	if err != nil {
		return convertError(op, err)
	}

	metadata := &auth.TokenMetadata{
		ClientId:  req.ClientId,
		AppId:     req.AppId,
		TokenType: "Bearer",
		Issuer:    "auth-service",
		Audiences: []string{"web-app"},
	}

	expiresAt := time.Now().Add(15 * time.Minute)

	return &auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(expiresAt),
		User:         models.ConvertUserToProto(userInfo),
		Metadata:     metadata,
	}, nil
}

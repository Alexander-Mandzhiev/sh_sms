package service

import (
	"backend/pkg/jwt_manager"
	client_apps "backend/protos/gen/go/apps/clients_apps"
	"backend/service/auth/handle"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *AuthService) IntrospectToken(ctx context.Context, token string, tokenTypeHint string) (*models.TokenIntrospection, error) {
	const op = "auth.service.IntrospectToken"
	logger := s.logger.With(slog.String("op", op), slog.String("token_type", tokenTypeHint))

	if token == "" {
		logger.Warn("empty token received")
		return nil, fmt.Errorf("%w: %v", handle.ErrInvalidToken, "empty token")
	}

	var tokenType jwt_manager.TokenType
	switch tokenTypeHint {
	case "refresh":
		tokenType = jwt_manager.RefreshToken
	default:
		tokenType = jwt_manager.AccessToken
	}

	session, err := s.session.GetSessionByToken(ctx, jwt_manager.HashToken(token))
	if err != nil {
		if errors.Is(err, handle.ErrSessionNotFound) {
			logger.Warn("session not found or inactive")
			return &models.TokenIntrospection{Active: false}, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	clientApp, err := s.clientApps.GetClientApp(ctx, &client_apps.IdentifierRequest{ClientId: session.ClientID.String(), AppId: int32(session.AppID)})
	if err != nil || clientApp == nil || !clientApp.IsActive {
		logger.Warn("client/app not active", slog.String("client_id", session.ClientID.String()), slog.Int("app_id", session.AppID))
		return &models.TokenIntrospection{Active: false}, nil
	}

	secret, err := s.getJWTSecret(ctx, session.ClientID, session.AppID, tokenType)
	if err != nil {
		logger.Error("failed to get JWT secret", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	claims, err := jwt_manager.Parse(token, secret, tokenType)
	if err != nil {
		logger.Warn("token parse failed", "error", err)
		return &models.TokenIntrospection{Active: false}, nil
	}

	return &models.TokenIntrospection{
		Active:      session.IsActive(),
		ClientID:    session.ClientID,
		UserID:      session.UserID,
		TokenType:   string(tokenType),
		ExpiresAt:   session.ExpiresAt,
		IssuedAt:    session.CreatedAt,
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
		Metadata: models.IntrospectMetadata{
			IPAddress: session.IPAddress.String(),
			UserAgent: session.UserAgent,
			ClientApp: fmt.Sprintf("%s:%d", session.ClientID, session.AppID),
			SessionID: session.SessionID.String(),
			AppID:     int32(session.AppID),
		},
	}, nil
}

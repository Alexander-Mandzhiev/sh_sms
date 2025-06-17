package service

import (
	"backend/pkg/jwt_manager"
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/users"
	"backend/service/auth/handle"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string, clientID uuid.UUID, appID int) (*models.UserInfo, string, string, error) {
	const op = "service.RefreshToken"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.Int("app_id", appID))
	logger.Debug("starting token refresh process")

	secret, err := s.getJWTSecret(ctx, clientID, appID, "refresh")
	if err != nil {
		logger.Error("failed to get JWT secret", slog.Any("error", err), slog.String("token_type", "refresh"))
		return nil, "", "", fmt.Errorf("secret retrieval failed: %w", err)
	}
	logger.Debug("JWT secret successfully retrieved")

	claims, err := jwt_manager.Parse(refreshToken, secret, jwt_manager.RefreshToken)
	if err != nil {
		logger.Warn("failed to parse refresh token", slog.Any("error", err), slog.String("token_hash", jwt_manager.HashToken(refreshToken)))
		return nil, "", "", fmt.Errorf("%w: invalid token", handle.ErrInvalidToken)
	}
	logger.Debug("token parsed successfully", slog.String("user_id", claims.UserID.String()))

	if claims.ClientID != clientID || claims.AppID != appID {
		logger.Warn("client/app mismatch", slog.String("token_client", claims.ClientID.String()), slog.Int("token_app", claims.AppID))
		return nil, "", "", handle.ErrPermissionDenied
	}

	session, err := s.findActiveSession(ctx, claims.UserID, clientID, appID, refreshToken)
	if err != nil {
		if errors.Is(err, handle.ErrSessionNotFound) {
			logger.Warn("active session not found", slog.String("user_id", claims.UserID.String()), slog.String("token_hash", jwt_manager.HashToken(refreshToken)))
		} else {
			logger.Error("session lookup failed", slog.Any("error", err), slog.String("user_id", claims.UserID.String()))
		}
		return nil, "", "", err
	}
	logger.Info("session found", slog.String("session_id", session.SessionID.String()), slog.Time("expires_at", session.ExpiresAt))

	user, err := s.users.GetUser(ctx, &users.GetRequest{Id: session.UserID.String(), ClientId: clientID.String()})
	if err != nil || !user.IsActive {
		logger.Warn("user not active or not found", slog.String("user_id", session.UserID.String()))
		return nil, "", "", fmt.Errorf("%w: user inactive", handle.ErrPermissionDenied)
	}
	logger.Debug("user validation passed", slog.String("user_email", user.Email))

	userId, err := utils.ValidateStringAndReturnUUID(user.Id)
	if err != nil {
		s.logger.Warn("Invalid user id", "client_id", clientID, "app_id", int(appID))
		return nil, "", "", err
	}

	rolesSet, err := s.getUserRoles(ctx, userId, clientID, appID)
	if err != nil {
		logger.Error("failed to get user roles", slog.Any("error", err), slog.String("user_id", user.Id))
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Debug("roles retrieved", slog.Int("count", len(rolesSet)))

	permissionSet, err := s.getUserPermissions(ctx, clientID, appID, rolesSet)
	if err != nil {
		logger.Error("failed to get permissions", slog.Any("error", err), slog.String("user_id", user.Id))
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Debug("permissions retrieved", slog.Int("count", len(permissionSet)))

	newAccess, err := s.generateTokens(ctx, session.UserID, clientID, appID, rolesSet, permissionSet, "access", s.cfg.JWT.AccessDuration)
	if err != nil {
		logger.Error("access token generation failed", slog.Any("error", err), slog.String("user_id", user.Id))
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	newRefresh, err := s.generateTokens(ctx, session.UserID, clientID, appID, rolesSet, permissionSet, "refresh", s.cfg.JWT.RefreshDuration)
	if err != nil {
		logger.Error("refresh token generation failed", slog.Any("error", err), slog.String("user_id", user.Id))
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("tokens generated successfully", slog.String("access_hash", jwt_manager.HashToken(newAccess)), slog.String("refresh_hash", jwt_manager.HashToken(newRefresh)))

	if err = s.UpdateSessionTokens(ctx, session.SessionID, newAccess, newRefresh); err != nil {
		logger.Error("failed to update session tokens", slog.Any("error", err), slog.String("session_id", session.SessionID.String()))
		return nil, "", "", err
	}
	logger.Info("session tokens updated", slog.String("session_id", session.SessionID.String()), slog.Time("new_expires_at", time.Now().Add(s.cfg.JWT.AccessDuration)))
	return &models.UserInfo{
		ID:          user.Id,
		Email:       user.Email,
		Phone:       user.Phone,
		FullName:    user.FullName,
		IsActive:    user.IsActive,
		Roles:       rolesSet,
		Permissions: permissionSet,
	}, newAccess, newRefresh, nil
}

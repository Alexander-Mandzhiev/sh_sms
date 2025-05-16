package service

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/apps/clients_apps"
	"backend/protos/gen/go/sso/users"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

const (
	accessTokenExpiry  = 15 * time.Minute
	refreshTokenExpiry = 7 * 24 * time.Hour
)

func (s *AuthService) Login(ctx context.Context, req models.AuthRequest) (*models.UserInfo, string, string, error) {
	const op = "auth.service.Login"
	s.logger.Debug("Login attempt", "client_id", req.ClientID, "app_id", req.AppID, "login", req.Login)

	clientID, err := utils.ValidateAndReturnUUID(req.ClientID)
	if err != nil {
		s.logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, "", "", err
	}

	clientApp, err := s.clientApps.GetClientApp(ctx, &client_apps.IdentifierRequest{ClientId: clientID.String(), AppId: int32(req.AppID)})
	if err != nil || clientApp == nil {
		s.logger.Warn("Invalid client/app combination", "client_id", req.ClientID, "app_id", req.AppID, "error", err)
		return nil, "", "", errors.New("invalid client or application")
	}

	userResp, err := s.users.GetUserByLogin(ctx, &users.GetUserByLoginRequest{
		Login:    req.Login,
		ClientId: clientID.String(),
		Password: req.Password,
	})
	if err != nil {
		s.logger.Error("Authentication failed", "error", err, "client_id", req.ClientID, "app_id", req.AppID)
		return nil, "", "", err
	}

	userId, err := utils.ValidateAndReturnUUID(userResp.Id)
	if err != nil {
		s.logger.Warn("Invalid user id", "client_id", req.ClientID, "app_id", req.AppID)
		return nil, "", "", err
	}

	rolesSet, err := s.getUserRoles(ctx, clientID, req.AppID)
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	permissionSet, err := s.getUserPermissions(ctx, clientID, req.AppID, rolesSet)
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err := s.generateTokens(ctx, userId, clientID, req.AppID, rolesSet, permissionSet, "access", accessTokenExpiry)
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, err := s.generateTokens(ctx, userId, clientID, req.AppID, rolesSet, permissionSet, "refresh", refreshTokenExpiry)
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	if err = s.createSession(ctx, userId, clientID, req.AppID, accessToken, refreshToken); err != nil {
		s.logger.Error("Session creation failed", "error", err, "user_id", userResp.Id)
		return nil, "", "", err
	}

	return &models.UserInfo{
		ID:          userResp.Id,
		Email:       userResp.Email,
		FullName:    userResp.FullName,
		IsActive:    userResp.IsActive,
		Roles:       rolesSet,
		Permissions: permissionSet,
	}, accessToken, refreshToken, nil
}

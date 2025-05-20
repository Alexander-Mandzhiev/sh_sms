package service

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/apps/app_manager"
	"backend/protos/gen/go/apps/clients_apps"
	"backend/protos/gen/go/sso/users"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log/slog"
)

func (s *AuthService) Login(ctx context.Context, req models.AuthRequest) (*models.UserInfo, string, string, error) {
	const op = "auth.service.Login"
	s.logger.Debug("Login attempt", "client_id", req.ClientID, "app_id", req.AppID, "login", req.Login)

	clientID, err := utils.ValidateAndReturnUUID(req.ClientID)
	if err != nil {
		s.logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, "", "", err
	}

	if err = utils.ValidateAppID(req.AppID); err != nil {
		s.logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, "", "", err
	}

	clientApp, err := s.clientApps.GetClientApp(ctx, &client_apps.IdentifierRequest{ClientId: clientID.String(), AppId: int32(req.AppID)})
	if err != nil || clientApp == nil {
		s.logger.Warn("Invalid client/app combination", "client_id", req.ClientID, "app_id", req.AppID, "error", err)
		return nil, "", "", errors.New("invalid client or application")
	}

	appID, err := s.appManager.GetApp(ctx, &app_manager.AppIdentifier{Identifier: &app_manager.AppIdentifier_Id{Id: int32(req.AppID)}})
	if err != nil {
		s.logger.Warn("Invalid app combination", "app_id", req.AppID, "error", err)
		return nil, "", "", err
	}
	userResp, err := s.users.GetUserByLogin(ctx, &users.GetUserByLoginRequest{
		Login:    req.Login,
		ClientId: clientID.String(),
		Password: req.Password,
	})
	if err != nil {
		s.logger.Error("Authentication failed", "error", err, "client_id", req.ClientID, "app_id", int(appID.Id))
		return nil, "", "", err
	}

	userId, err := utils.ValidateAndReturnUUID(userResp.Id)
	if err != nil {
		s.logger.Warn("Invalid user id", "client_id", req.ClientID, "app_id", int(appID.Id))
		return nil, "", "", err
	}

	rolesSet, err := s.getUserRoles(ctx, userId, clientID, int(appID.Id))
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	permissionSet, err := s.getUserPermissions(ctx, clientID, int(appID.Id), rolesSet)
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	accessToken, err := s.generateTokens(ctx, userId, clientID, int(appID.Id), rolesSet, permissionSet, "access", s.cfg.JWT.AccessDuration)
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, err := s.generateTokens(ctx, userId, clientID, int(appID.Id), rolesSet, permissionSet, "refresh", s.cfg.JWT.RefreshDuration)
	if err != nil {
		return nil, "", "", fmt.Errorf("%s: %w", op, err)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	var userAgent, ipAddress string

	if ok {
		if ua := md.Get("user-agent"); len(ua) > 0 {
			userAgent = ua[0]
		}

		ipHeaders := []string{"x-forwarded-for", "x-real-ip"}
		for _, header := range ipHeaders {
			if ip := md.Get(header); len(ip) > 0 {
				ipAddress = ip[0]
				break
			}
		}

	}

	if ipAddress == "" {
		ipAddress = "0.0.0.0"
	}
	if err = s.CreateSession(ctx, userId, clientID, int(appID.Id), accessToken, refreshToken, ipAddress, userAgent); err != nil {
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

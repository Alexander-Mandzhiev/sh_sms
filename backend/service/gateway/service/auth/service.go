package auth_service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/auth"
)

type AuthService interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error)
	Logout(ctx context.Context, req *auth.LogoutRequest) error
	RefreshToken(ctx context.Context, req *auth.RefreshRequest) (*auth.AuthResponse, error)
	ValidateToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenInfo, error)
	IntrospectToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenIntrospection, error)
	CheckPermission(ctx context.Context, req *auth.PermissionCheckRequest) (*auth.PermissionCheckResponse, error)
	ListSessionsForUser(ctx context.Context, req *auth.SessionFilter) (*auth.SessionList, error)
	ListAllSessions(ctx context.Context, req *auth.AllSessionsFilter) (*auth.SessionList, error)
	TerminateSession(ctx context.Context, req *auth.SessionID) error
}

type authService struct {
	client auth.AuthServiceClient

	logger *slog.Logger
}

func NewAuthService(client auth.AuthServiceClient, logger *slog.Logger) AuthService {
	return &authService{
		client: client,
		logger: logger.With("service", "auth"),
	}
}

func (s *authService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	const op = "authService.Login"
	logger := s.logger.With(slog.String("op", op))

	res, err := s.client.Login(ctx, req)
	if err != nil {
		logger.Error("gRPC Login failed", "error", err)
		return nil, err
	}

	logger.Info("successful login", "client_id", req.ClientId)
	return res, nil
}

func (s *authService) Logout(ctx context.Context, req *auth.LogoutRequest) error {
	const op = "authService.Logout"
	logger := s.logger.With(slog.String("op", op))

	_, err := s.client.Logout(ctx, req)
	if err != nil {
		logger.Error("gRPC Logout failed", "error", err)
		return err
	}

	logger.Info("successful logout")
	return nil
}

func (s *authService) RefreshToken(ctx context.Context, req *auth.RefreshRequest) (*auth.AuthResponse, error) {
	const op = "authService.RefreshToken"
	logger := s.logger.With(slog.String("op", op))

	res, err := s.client.RefreshToken(ctx, req)
	if err != nil {
		logger.Error("gRPC RefreshToken failed", "error", err)
		return nil, err
	}

	logger.Info("token refreshed", "client_id", req.ClientId)
	return res, nil
}

func (s *authService) ValidateToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenInfo, error) {
	const op = "authService.ValidateToken"
	logger := s.logger.With(slog.String("op", op))

	res, err := s.client.ValidateToken(ctx, req)
	if err != nil {
		logger.Error("gRPC ValidateToken failed", "error", err)
		return nil, err
	}

	logger.Debug("token validated", "active", res.Active)
	return res, nil
}

func (s *authService) IntrospectToken(ctx context.Context, req *auth.ValidateRequest) (*auth.TokenIntrospection, error) {
	const op = "authService.IntrospectToken"
	logger := s.logger.With(slog.String("op", op))

	res, err := s.client.IntrospectToken(ctx, req)
	if err != nil {
		logger.Error("gRPC IntrospectToken failed", "error", err)
		return nil, err
	}

	logger.Debug("token introspected", "active", res.Active)
	return res, nil
}

func (s *authService) CheckPermission(ctx context.Context, req *auth.PermissionCheckRequest) (*auth.PermissionCheckResponse, error) {
	const op = "authService.CheckPermission"
	logger := s.logger.With(slog.String("op", op))

	res, err := s.client.CheckPermission(ctx, req)
	if err != nil {
		logger.Error("gRPC CheckPermission failed", "error", err)
		return nil, err
	}

	logger.Info("permission checked", "allowed", res.Allowed)
	return res, nil
}

func (s *authService) ListSessionsForUser(ctx context.Context, req *auth.SessionFilter) (*auth.SessionList, error) {
	const op = "authService.ListSessionsForUser"
	logger := s.logger.With(slog.String("op", op))

	res, err := s.client.ListSessionsForUser(ctx, req)
	if err != nil {
		logger.Error("gRPC ListSessionsForUser failed", "error", err)
		return nil, err
	}

	logger.Info("sessions listed", "count", len(res.Sessions))
	return res, nil
}

func (s *authService) ListAllSessions(ctx context.Context, req *auth.AllSessionsFilter) (*auth.SessionList, error) {
	const op = "authService.ListAllSessions"
	logger := s.logger.With(slog.String("op", op))

	res, err := s.client.ListAllSessions(ctx, req)
	if err != nil {
		logger.Error("gRPC ListAllSessions failed", "error", err)
		return nil, err
	}

	logger.Info("all sessions listed", "count", len(res.Sessions))
	return res, nil
}

func (s *authService) TerminateSession(ctx context.Context, req *auth.SessionID) error {
	const op = "authService.TerminateSession"
	logger := s.logger.With(slog.String("op", op))

	_, err := s.client.TerminateSession(ctx, req)
	if err != nil {
		logger.Error("gRPC TerminateSession failed", "error", err)
		return err
	}

	logger.Info("session terminated", "session_id", req.SessionId)
	return nil
}

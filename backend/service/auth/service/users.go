package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/users"
)

type UserService interface {
	GetUser(ctx context.Context, req *users.GetRequest) (*users.User, error)
	GetUserByLogin(ctx context.Context, req *users.GetUserByLoginRequest) (*users.UserInfo, error)
	BatchGetUsers(ctx context.Context, req *users.BatchGetRequest) (*users.BatchGetResponse, error)
}

type userService struct {
	client users.UserServiceClient
	logger *slog.Logger
}

func NewUserService(client users.UserServiceClient, logger *slog.Logger) UserService {
	return &userService{
		client: client,
		logger: logger.With("service", "user"),
	}
}

func (s *userService) GetUser(ctx context.Context, req *users.GetRequest) (*users.User, error) {
	s.logger.Debug("getting user", "user_id", req.Id, "client_id", req.ClientId)
	return s.client.GetUser(ctx, req)
}

func (s *userService) GetUserByLogin(ctx context.Context, req *users.GetUserByLoginRequest) (*users.UserInfo, error) {
	s.logger.Debug("get user by login", "user_id", req.Login)
	return s.client.GetUserByLogin(ctx, req)
}
func (s *userService) BatchGetUsers(ctx context.Context, req *users.BatchGetRequest) (*users.BatchGetResponse, error) {
	s.logger.Debug("batch getting users", "client_id", req.ClientId, "user_count", len(req.UserIds), "include_inactive", req.IncludeInactive)

	resp, err := s.client.BatchGetUsers(ctx, req)
	if err != nil {
		s.logger.Error("failed to batch get users", "error", err, "client_id", req.ClientId, "user_count", len(req.UserIds))
		return nil, err
	}

	s.logger.Debug("batch get users completed", "found", len(resp.Users), "missing", len(resp.MissingIds))
	return resp, nil
}

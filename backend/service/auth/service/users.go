package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/users"
)

type UserService interface {
	CreateUser(ctx context.Context, req *users.CreateRequest) (*users.User, error)
	GetUser(ctx context.Context, req *users.GetRequest) (*users.User, error)
	UpdateUser(ctx context.Context, req *users.UpdateRequest) (*users.User, error)
	DeleteUser(ctx context.Context, req *users.DeleteRequest) (*users.SuccessResponse, error)
	ListUsers(ctx context.Context, req *users.ListRequest) (*users.ListResponse, error)
	SetPassword(ctx context.Context, req *users.SetPasswordRequest) (*users.SuccessResponse, error)
	RestoreUser(ctx context.Context, req *users.RestoreRequest) (*users.User, error)
	GetUserByLogin(ctx context.Context, req *users.GetUserByLoginRequest) (*users.UserInfo, error)
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

func (s *userService) CreateUser(ctx context.Context, req *users.CreateRequest) (*users.User, error) {
	s.logger.Debug("creating user", "email", req.Email)
	return s.client.CreateUser(ctx, req)
}

func (s *userService) GetUser(ctx context.Context, req *users.GetRequest) (*users.User, error) {
	s.logger.Debug("getting user", "user_id", req.Id, "client_id", req.ClientId)
	return s.client.GetUser(ctx, req)
}

func (s *userService) UpdateUser(ctx context.Context, req *users.UpdateRequest) (*users.User, error) {
	s.logger.Debug("updating user", "user_id", req.Id, "client_id", req.ClientId)
	return s.client.UpdateUser(ctx, req)
}

func (s *userService) DeleteUser(ctx context.Context, req *users.DeleteRequest) (*users.SuccessResponse, error) {
	s.logger.Debug("deleting user", "user_id", req.Id, "client_id", req.ClientId)
	return s.client.DeleteUser(ctx, req)
}

func (s *userService) ListUsers(ctx context.Context, req *users.ListRequest) (*users.ListResponse, error) {
	s.logger.Debug("listing users", "client_id", req.ClientId)
	return s.client.ListUsers(ctx, req)
}

func (s *userService) SetPassword(ctx context.Context, req *users.SetPasswordRequest) (*users.SuccessResponse, error) {
	s.logger.Debug("setting password", "user_id", req.Id, "client_id", req.ClientId)
	return s.client.SetPassword(ctx, req)
}

func (s *userService) RestoreUser(ctx context.Context, req *users.RestoreRequest) (*users.User, error) {
	s.logger.Debug("restoring user", "user_id", req.Id, "client_id", req.ClientId)
	return s.client.RestoreUser(ctx, req)
}

func (s *userService) GetUserByLogin(ctx context.Context, req *users.GetUserByLoginRequest) (*users.UserInfo, error) {
	s.logger.Debug("get user by login", "user_id", req.Login)
	return s.client.GetUserByLogin(ctx, req)
}

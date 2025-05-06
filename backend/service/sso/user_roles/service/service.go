package service

import (
	"backend/service/sso/models"
	serviceRole "backend/service/sso/roles/service"
	serviceUser "backend/service/sso/users/service"
	"context"
	"github.com/google/uuid"

	"errors"
	"log/slog"
)

var (
	ErrInternal           = errors.New("internal server error")
	ErrAssignmentExists   = errors.New("assignment already exists")
	ErrRoleNotFound       = errors.New("role not found")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrUserNotFound       = errors.New("user not found")
	ErrAssignmentNotFound = errors.New("assignment not found")
)

type UserRoleProvider interface {
	Assign(ctx context.Context, role *models.UserRole) (*models.UserRole, error)
	Revoke(ctx context.Context, userID uuid.UUID, roleID uuid.UUID, clientID uuid.UUID, appID int) error
	Exists(ctx context.Context, userID, roleID, clientID uuid.UUID, appID int) (bool, error)
	ListForUser(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error)
	ListForRole(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error)
}

type Service struct {
	logger       *slog.Logger
	provider     UserRoleProvider
	roleProvider serviceRole.RolesProvider
	userProvider serviceUser.UsersProvider
}

func New(provider UserRoleProvider, roleProvider serviceRole.RolesProvider, userProvider serviceUser.UsersProvider, logger *slog.Logger) *Service {
	const op = "service.New.UserRole"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing sso service - handle user_role", slog.String("op", op))
	return &Service{
		provider:     provider,
		roleProvider: roleProvider,
		userProvider: userProvider,
		logger:       logger,
	}
}

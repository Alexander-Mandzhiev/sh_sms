package groups_service

import (
	"backend/pkg/models/groups"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type GroupsProvider interface {
	CreateGroup(ctx context.Context, s *groups_models.CreateGroup) (*groups_models.Group, error)
	GetGroup(ctx context.Context, publicID, clientID uuid.UUID) (*groups_models.Group, error)
	ListGroups(ctx context.Context, params *groups_models.ListGroupsRequest) (*groups_models.GroupListResponse, error)
	UpdateGroup(ctx context.Context, ug *groups_models.UpdateGroup) (*groups_models.Group, error)
	DeleteGroup(ctx context.Context, publicID, clientID uuid.UUID) error
}

type Service struct {
	logger   *slog.Logger
	provider GroupsProvider
}

func New(provider GroupsProvider, logger *slog.Logger) *Service {
	const op = "service.New.PrivateSchool.Groups"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service groups", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}

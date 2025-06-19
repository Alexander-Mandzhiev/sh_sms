package groups_handle

import (
	groups_models "backend/pkg/models/groups"
	"backend/protos/gen/go/private_school"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type GroupsService interface {
	CreateGroup(ctx context.Context, req *groups_models.CreateGroup) (*groups_models.Group, error)
	GetGroup(ctx context.Context, publicID, clientID uuid.UUID) (*groups_models.Group, error)
	ListGroups(ctx context.Context, req *groups_models.ListGroupsRequest) (*groups_models.GroupListResponse, error)
	UpdateGroup(ctx context.Context, req *groups_models.UpdateGroup) (*groups_models.Group, error)
	DeleteGroup(ctx context.Context, publicID, clientID uuid.UUID) error
}

type serverAPI struct {
	private_school_v1.UnimplementedGroupServiceServer
	service GroupsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service GroupsService, logger *slog.Logger) {
	private_school_v1.RegisterGroupServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "groups"),
	})
}

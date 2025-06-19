package groups_handle

import (
	"backend/pkg/models/groups"
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

type ServerAPI struct {
	private_school_v1.UnimplementedGroupServiceServer
	Service GroupsService
	Logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service GroupsService, logger *slog.Logger) {
	private_school_v1.RegisterGroupServiceServer(gRPCServer, &ServerAPI{
		Service: service,
		Logger:  logger.With("module", "groups"),
	})
}

func NewServerAPI(service GroupsService, logger *slog.Logger) *ServerAPI {
	return &ServerAPI{
		Service: service,
		Logger:  logger,
	}
}

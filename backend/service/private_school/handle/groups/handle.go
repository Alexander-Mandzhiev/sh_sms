package groups_handle

import (
	"backend/protos/gen/go/private_school"
	"google.golang.org/grpc"
	"log/slog"
)

type GroupsService interface {
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

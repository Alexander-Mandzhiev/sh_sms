package subjects_handle

import (
	"backend/protos/gen/go/library/subjects"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type SubjectsService interface {
	Create(ctx context.Context, req *subjects.CreateSubjectRequest) (*subjects.Subject, error)
	Get(ctx context.Context, req *subjects.GetSubjectRequest) (*subjects.Subject, error)
	Update(ctx context.Context, req *subjects.UpdateSubjectRequest) (*subjects.Subject, error)
	Delete(ctx context.Context, req *subjects.DeleteSubjectRequest) error
	List(ctx context.Context, req *subjects.ListSubjectsRequest) (*subjects.ListSubjectsResponse, error)
}

type serverAPI struct {
	subjects.UnimplementedSubjectsServiceServer
	service SubjectsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service SubjectsService, logger *slog.Logger) {
	subjects.RegisterSubjectsServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "subjects"),
	})
}

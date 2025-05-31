package subjects_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"

	"google.golang.org/grpc"
	"log/slog"
)

type SubjectsService interface {
	CreateSubject(ctx context.Context, params *library_models.CreateSubjectParams) (*library_models.Subject, error)
	GetSubject(ctx context.Context, id int32) (*library_models.Subject, error)
	UpdateSubject(ctx context.Context, params *library_models.UpdateSubjectParams) (*library_models.Subject, error)
	DeleteSubject(ctx context.Context, id int32) error
	ListSubjects(ctx context.Context) ([]*library_models.Subject, error)
}

type serverAPI struct {
	library.UnimplementedSubjectServiceServer
	service SubjectsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service SubjectsService, logger *slog.Logger) {
	library.RegisterSubjectServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "subjects"),
	})
}

package subjects_handle

import (
	private_school_models "backend/pkg/models/private_school"
	library "backend/protos/gen/go/library"
	"context"

	"google.golang.org/grpc"
	"log/slog"
)

type SubjectsService interface {
	CreateSubject(ctx context.Context, params *private_school_models.CreateSubjectParams) (*private_school_models.Subject, error)
	GetSubject(ctx context.Context, id int32) (*private_school_models.Subject, error)
	UpdateSubject(ctx context.Context, params *private_school_models.UpdateSubjectParams) (*private_school_models.Subject, error)
	DeleteSubject(ctx context.Context, id int32) error
	ListSubjects(ctx context.Context) ([]*private_school_models.Subject, error)
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

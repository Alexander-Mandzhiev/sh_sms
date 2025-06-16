package teachers_handle

import (
	"backend/pkg/models/private_school"
	"backend/protos/gen/go/private_school"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type TeacherService interface {
	CreateTeacher(ctx context.Context, teacher *private_school_models.CreateTeacher) (*private_school_models.Teacher, error)
	GetTeacher(ctx context.Context, id, clientID uuid.UUID) (*private_school_models.Teacher, error)
	ListTeachers(ctx context.Context, filter *private_school_models.ListTeachersFilter) (*private_school_models.ListTeachersResponse, error)
	UpdateTeacher(ctx context.Context, update *private_school_models.UpdateTeacher) (*private_school_models.Teacher, error)
	DeleteTeacher(ctx context.Context, id, clientID uuid.UUID) error
}
type serverAPI struct {
	private_school_v1.UnimplementedTeacherServiceServer
	service TeacherService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service TeacherService, logger *slog.Logger) {
	private_school_v1.RegisterTeacherServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "teachers"),
	})
}

package teachers_handle

import (
	"backend/pkg/models/teacher"
	"backend/protos/gen/go/private_school"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type TeacherService interface {
	CreateTeacher(ctx context.Context, teacher *teachers_models.CreateTeacher) (*teachers_models.Teacher, error)
	GetTeacher(ctx context.Context, id, clientID uuid.UUID) (*teachers_models.Teacher, error)
	ListTeachers(ctx context.Context, filter *teachers_models.ListTeachersFilter) (*teachers_models.ListTeachersResponse, error)
	UpdateTeacher(ctx context.Context, update *teachers_models.UpdateTeacher) (*teachers_models.Teacher, error)
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

package students_handle

import (
	"backend/pkg/models/students"
	"backend/protos/gen/go/private_school"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student *students_models.CreateStudent) (*students_models.Student, error)
	UpdateStudent(ctx context.Context, updateData *students_models.UpdateStudent) (*students_models.Student, error)

	GetStudent(ctx context.Context, id, clientID uuid.UUID) (*students_models.Student, error)
	ListStudents(ctx context.Context, params *students_models.ListStudentsRequest) ([]*students_models.Student, string, error)

	HardDeleteStudent(ctx context.Context, id, clientID uuid.UUID) error
	SoftDeleteStudent(ctx context.Context, id, clientID uuid.UUID) error
	RestoreStudent(ctx context.Context, id, clientID uuid.UUID) error
}

type serverAPI struct {
	private_school_v1.UnimplementedStudentServiceServer
	service StudentService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service StudentService, logger *slog.Logger) {
	private_school_v1.RegisterStudentServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "students"),
	})
}

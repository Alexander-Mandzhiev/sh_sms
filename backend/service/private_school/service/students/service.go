package students_service

import (
	students_models "backend/pkg/models/students"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type StudentsProvider interface {
	CreateStudent(ctx context.Context, student *students_models.CreateStudent) (*students_models.Student, error)
	UpdateStudent(ctx context.Context, update *students_models.UpdateStudent) (*students_models.Student, error)
	GetStudent(ctx context.Context, id, clientID uuid.UUID) (*students_models.Student, error)
	ListStudents(ctx context.Context, params *students_models.ListStudentsRequest) ([]*students_models.Student, *students_models.Cursor, error)

	HardDeleteStudent(ctx context.Context, id, clientID uuid.UUID) error
	SoftDeleteStudent(ctx context.Context, id, clientID uuid.UUID) error
	RestoreStudent(ctx context.Context, id, clientID uuid.UUID) error
}

type Service struct {
	logger   *slog.Logger
	provider StudentsProvider
}

func New(provider StudentsProvider, logger *slog.Logger) *Service {
	const op = "service.New.PrivateSchool.Students"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service students", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
